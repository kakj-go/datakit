package hostobject

import (
	"context"
	"encoding/json"
	"runtime"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/git"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var moduleLogger *logger.Logger

type objCollector struct {
	Name  string //deprecated
	Class string //deprecated
	Desc  string `toml:"description,omitempty"` //deprecated

	Interval datakit.Duration
	Pipeline string            `toml:"pipeline"`
	Tags     map[string]string `toml:"tags,omitempty"`

	p *pipeline.Pipeline

	ctx       context.Context
	cancelFun context.CancelFunc

	mode string

	testError error
}

func (c *objCollector) isTestOnce() bool {
	return c.mode == "test"
}

func (c *objCollector) isDebug() bool {
	return c.mode == "debug"
}

func (_ *objCollector) Catalog() string {
	return InputCat
}

func (_ *objCollector) SampleConfig() string {
	return SampleConfig
}

func (r *objCollector) PipelineConfig() map[string]string {
	return map[string]string{
		InputName: pipelineSample,
	}
}

func (c *objCollector) Run() {

	moduleLogger = logger.SLogger(InputName)

	if c.Interval.Duration == 0 {
		c.Interval.Duration = 5 * time.Minute
	}

	if c.Interval.Duration == 0 {
		c.Interval.Duration = 5 * time.Minute
	}

	defer func() {
		if e := recover(); e != nil {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				moduleLogger.Errorf("panic: %s", err)
				moduleLogger.Errorf("%s", string(buf[:n]))
			}
		}
	}()

	go func() {
		<-datakit.Exit.Wait()
		c.cancelFun()
	}()

	c.p = c.getPipeline()

	for {

		select {
		case <-c.ctx.Done():
			return
		default:
		}

		message := getHostObjectMessage()

		messageData, err := json.Marshal(message)
		if err != nil {
			moduleLogger.Errorf("json marshal err:%s", err.Error())
			datakit.SleepContext(c.ctx, c.Interval.Duration)
			continue
		}

		moduleLogger.Debugf("%s", string(messageData))

		fields := map[string]interface{}{
			"message":          string(messageData),
			"os":               message.Host.HostMeta.OS,
			"start_time":       message.Host.HostMeta.BootTime,
			"datakit_ver":      git.Version,
			"cpu_usage":        message.Host.cpuPercent,
			"mem_used_percent": message.Host.Mem.usedPercent,
			"load":             message.Host.load5,
			"state":            "online",
		}

		for k, v := range c.Tags {
			if _, ok := fields[k]; !ok {
				fields[k] = v
			}
		}

		if c.p != nil {
			if result, err := c.p.Run(string(messageData)).Result(); err == nil {
				moduleLogger.Debugf("%s", result)
				for k, v := range result {
					fields[k] = v
				}
			} else {
				moduleLogger.Errorf("%s", err)
			}
		}

		tags := map[string]string{
			"name": message.Host.HostMeta.HostName,
		}

		tm := time.Now().UTC()

		if c.isTestOnce() {
			// pass
		} else {
			io.NamedFeedEx(InputName, io.Object, "HOST", tags, fields, tm)
		}

		datakit.SleepContext(c.ctx, c.Interval.Duration)
	}
}

func (c *objCollector) getPipeline() *pipeline.Pipeline {

	fname := c.Pipeline
	if fname == "" {
		fname = InputName + ".p"
	}

	p, err := pipeline.NewPipelineByScriptPath(fname)
	if err != nil {
		moduleLogger.Warnf("%s", err)
		return nil
	}

	return p
}

func newInput(mode string) *objCollector {
	o := &objCollector{}
	o.mode = mode
	o.ctx, o.cancelFun = context.WithCancel(context.Background())
	return o
}

func init() {
	inputs.Add(InputName, func() inputs.Input {
		return newInput("")
	})
}
