package inputs

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/system/rtpanic"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
)

var (
	Inputs     = map[string]Creator{}
	InputsInfo = map[string][]*inputInfo{}

	l           = logger.DefaultSLogger("inputs")
	panicInputs = map[string]int{}
	mtx         = sync.RWMutex{}
)

func GetElectionInputs() []ElectionInput {
	res := []ElectionInput{}
	for k, arr := range InputsInfo {
		for _, x := range arr {
			switch y := x.input.(type) {
			case ElectionInput:
				l.Debugf("find election inputs %s", k)
				res = append(res, y)
			}
		}
	}
	return res
}

type Input interface {
	Catalog() string
	Run()
	SampleConfig() string
	// add more...
}

type HTTPInput interface {
	//Input
	RegHttpHandler()
}

type PipelineInput interface {
	//Input
	PipelineConfig() map[string]string
	RunPipeline()
}

// new input interface got extra interfaces, for better documentation
type InputV2 interface {
	Input
	SampleMeasurement() []Measurement
	AvailableArchs() []string
}

type ElectionInput interface {
	Pause() error
	Resume() error
}

type Creator func() Input

func Add(name string, creator Creator) {
	if _, ok := Inputs[name]; ok {
		l.Fatalf("inputs %s exist(from datakit)", name)
	}

	Inputs[name] = creator
}

type inputInfo struct {
	input Input
}

func (ii *inputInfo) Run() {
	if ii.input == nil {
		return
	}

	switch ii.input.(type) {
	case Input:
		ii.input.Run()
	default:
		l.Errorf("invalid input type")
	}
}

func AddInput(name string, input Input) error {
	mtx.Lock()
	defer mtx.Unlock()

	InputsInfo[name] = append(InputsInfo[name], &inputInfo{input: input})

	l.Debugf("add input %s, total %d", name, len(InputsInfo[name]))
	return nil
}

func AddSelf() {
	self, _ := Inputs["self"]
	AddInput("self", self())
}

func ResetInputs() {
	mtx.Lock()
	defer mtx.Unlock()
	InputsInfo = map[string][]*inputInfo{}
}

func RunInputs() error {
	mtx.RLock()
	defer mtx.RUnlock()

	for name, arr := range InputsInfo {
		for _, ii := range arr {
			if ii.input == nil {
				l.Debugf("skip non-datakit-input %s", name)
				continue
			}

			if inp, ok := ii.input.(HTTPInput); ok {
				inp.RegHttpHandler()
			}
			if inp, ok := ii.input.(PipelineInput); ok {
				inp.RunPipeline()
			}

			datakit.WG.Add(1)
			go func(name string, ii *inputInfo) {

				// NOTE: 让每个采集器间歇运行，防止每个采集器扎堆启动，导致主机资源消耗出现规律性的峰值
				time.Sleep(time.Duration(rand.Int63n(int64(10 * time.Second))))
				l.Infof("starting input %s ...", name)

				defer datakit.WG.Done()
				protectRunningInput(name, ii)
				l.Infof("input %s exited", name)
			}(name, ii)
		}
	}
	return nil
}

var (
	MaxCrash = 6
)

func protectRunningInput(name string, ii *inputInfo) {
	var f rtpanic.RecoverCallback
	crashTime := []string{}

	f = func(trace []byte, err error) {

		defer rtpanic.Recover(f, nil)

		if trace != nil {
			l.Warnf("input %s panic err: %v", name, err)
			l.Warnf("input %s panic trace:\n%s", name, string(trace))

			crashTime = append(crashTime, fmt.Sprintf("%v", time.Now()))
			addPanic(name)

			if len(crashTime) >= MaxCrash {
				l.Warnf("input %s crash %d times(at %+#v), exit now.",
					name, len(crashTime), strings.Join(crashTime, "\n"))
				return
			}
		}

		ii.Run()
	}

	f(nil, nil)
}

func GetPanicCnt(name string) int {
	mtx.RLock()
	defer mtx.RUnlock()

	return panicInputs[name]
}

func addPanic(name string) {
	mtx.Lock()
	defer mtx.Unlock()

	panicInputs[name]++
}

func InputEnabled(name string) (n int) {
	mtx.RLock()
	defer mtx.RUnlock()
	arr, ok := InputsInfo[name]
	if !ok {
		return
	}

	n = len(arr)
	l.Debugf("name enabled %d", n)
	return
}

func Init() {
	l = logger.SLogger("inputs")
}
