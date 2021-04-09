package rabbitmq

import (
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
	"time"
	"path/filepath"
)

func (_ *Input) SampleConfig() string {
	return sample
}

func (_ *Input) Catalog() string {
	return inputName
}

func (_ *Input) PipelineConfig() map[string]string {
	pipelineMap := map[string]string{
		"rabbitmq":           pipelineCfg,
	}
	return pipelineMap
}


func (n *Input) Run() {
	l = logger.SLogger(inputName)
	l.Info("rabbitmq start")

	if n.TailF != nil {
		go func() {
			if err := n.TailF.Init(); err != nil {
				l.Errorf("nginx init tailf err:%s", err.Error())
				return
			}
			if n.TailF.Option.Pipeline != "" {
				n.TailF.Option.Pipeline = filepath.Join(datakit.PipelineDir, n.TailF.Option.Pipeline)
			}else {
				n.TailF.Option.Pipeline = filepath.Join(datakit.PipelineDir, "nginx.p")
			}

			n.TailF.Run()
		}()
	}

	client, err := n.createHttpClient()
	if err != nil {
		l.Errorf("[error] rabbitmq init client err:%s", err.Error())
		return
	}
	n.client = client
	if n.Interval.Duration == 0 {
		n.Interval.Duration = time.Second * 30
	}

	tick := time.NewTicker(n.Interval.Duration)
	defer tick.Stop()
	cleanCacheTick := time.NewTicker(time.Second * 5)
	defer cleanCacheTick.Stop()

	for {
		select {
		case <-tick.C:
			n.getMetric()
		case <-cleanCacheTick.C:
			if len(collectCache) > 0 {
				err := inputs.FeedMeasurement(inputName, io.Metric, collectCache, &io.Option{CollectCost: time.Since(n.start)})
				collectCache = collectCache[:]
				if err != nil {
					l.Errorf(err.Error())
					continue
				}
			}
		case <-datakit.Exit.Wait():
			l.Info("rabbitmq exit")
			return
		}
	}
}

type MetricFunc func(n *Input)

func (n *Input) getMetric() {
	getFunc := []MetricFunc{getOverview, getNode, getQueues, getExchange}
	n.wg.Add(len(getFunc))
	for _, v := range getFunc {
		go func(gf MetricFunc) {
			defer n.wg.Done()
			gf(n)
		}(v)
	}
	n.wg.Wait()
}

func (n *Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{
		&OverviewMeasurement{},
		&QueueMeasurement{},
		&ExchangeMeasurement{},
		&NodeMeasurement{},
	}
}

func init() {
	inputs.Add(inputName, func() inputs.Input {
		s := &Input{}
		return s
	})
}
