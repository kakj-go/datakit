// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package opentelemetry is input for opentelemetry

package opentelemetry

import (
	"strings"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	dkHTTP "gitlab.jiagouyun.com/cloudcare-tools/datakit/http"
	itrace "gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/trace"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/opentelemetry/collector"
)

var (
	_ inputs.InputV2   = &Input{}
	_ inputs.HTTPInput = &Input{}
)

const (
	inputName    = "opentelemetry"
	sampleConfig = `
[[inputs.opentelemetry]]
  ## 在创建'trace',Span','resource'时，会加入很多标签，这些标签最终都会出现在'Span'中
  ## 当您不希望这些标签太多造成网络上不必要的流量损失时，可选择忽略掉这些标签
  ## 支持正则表达，注意:将所有的'.'替换成'_'
  ## When creating 'trace', 'span' and 'resource', many labels will be added, and these labels will eventually appear in all 'spans'
  ## When you don't want too many labels to cause unnecessary traffic loss on the network, you can choose to ignore these labels
  ## Support regular expression. Note!!!: all '.' Replace with '_'
  # ignore_attribute_keys = ["os_*","process_*"]

  ## Keep rare tracing resources list switch.
  ## If some resources are rare enough(not presend in 1 hour), those resource will always send
  ## to data center and do not consider samplers and filters.
  # keep_rare_resource = false

  ## By default every error presents in span will be send to data center and omit any filters or
  ## sampler. If you want to get rid of some error status, you can set the error status list here.
  # omit_err_status = ["404"]

  ## Ignore tracing resources map like service:[resources...].
  ## The service name is the full service name in current application.
  ## The resource list is regular expressions uses to block resource names.
  ## If you want to block some resources universally under all services, you can set the
  ## service name as "*". Note: double quotes "" cannot be omitted.
  # [inputs.opentelemetry.close_resource]
    # service1 = ["resource1", "resource2", ...]
    # service2 = ["resource1", "resource2", ...]
    # "*" = ["close_resource_under_all_services"]
    # ...

  ## Sampler config uses to set global sampling strategy.
  ## sampling_rate used to set global sampling rate.
  # [inputs.opentelemetry.sampler]
    # sampling_rate = 1.0

  # [inputs.opentelemetry.tags]
    # key1 = "value1"
    # key2 = "value2"
    # ...

  [inputs.opentelemetry.expectedHeaders]
    ## 如有header配置 则请求中必须要携带 否则返回状态码500
  ## 可作为安全检测使用,必须全部小写
  # ex_version = xxx
  # ex_name = xxx
  # ...

  ## grpc
  [inputs.opentelemetry.grpc]
  ## trace for grpc
  trace_enable = false

  ## metric for grpc
  metric_enable = false

  ## grpc listen addr
  addr = "127.0.0.1:9550"

  ## http
  [inputs.opentelemetry.http]
  ## if enable=true
  ## http path (do not edit):
  ##	trace : /otel/v1/trace
  ##	metric: /otel/v1/metric
  ## use as : http://127.0.0.1:9529/otel/v1/trace . Method = POST
  enable = true
  ## return to client status_ok_code :200/202
  http_status_ok = 200
`
)

var (
	l       = logger.DefaultSLogger("otel-log")
	sampler *itrace.Sampler
)

type Input struct {
	Pipelines           map[string]string   `toml:"pipelines"` // deprecated
	Ogrpc               *otlpGrpcCollector  `toml:"grpc"`
	OHTTPc              *otlpHTTPCollector  `toml:"http"`
	CloseResource       map[string][]string `toml:"close_resource"`
	OmitErrStatus       []string            `toml:"omit_err_status"`
	Sampler             *itrace.Sampler     `toml:"sampler"`
	IgnoreAttributeKeys []string            `toml:"ignore_attribute_keys"`
	Tags                map[string]string   `toml:"tags"`
	ExpectedHeaders     map[string]string   `toml:"expectedHeaders"`
	inputName           string
	semStop             *cliutils.Sem // start stop signal
}

func (*Input) Catalog() string {
	return inputName
}

func (*Input) AvailableArchs() []string {
	return datakit.AllOS
}

func (*Input) SampleConfig() string {
	return sampleConfig
}

func (*Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{&itrace.TraceMeasurement{Name: inputName}}
}

func (ipt *Input) RegHTTPHandler() {
	dkHTTP.RegHTTPHandler("POST", "/otel/v1/trace", ipt.OHTTPc.apiOtlpTrace)
	dkHTTP.RegHTTPHandler("POST", "/otel/v1/metric", ipt.OHTTPc.apiOtlpMetric)
}

func (ipt *Input) exit() {
	ipt.Ogrpc.stop()
}

func (ipt *Input) Terminate() {
	if ipt.semStop != nil {
		ipt.semStop.Close()
	}
}

func (ipt *Input) Run() {
	l = logger.SLogger("otlp-log")
	storage := collector.NewSpansStorage()
	// add filters: the order of appending filters into AfterGather is important!!!
	// the order of appending represents the order of that filter executes.
	// add close resource filter
	if len(ipt.CloseResource) != 0 {
		closeResource := &itrace.CloseResource{}
		closeResource.UpdateIgnResList(ipt.CloseResource)
		storage.AfterGather.AppendFilter(closeResource.Close)
	}
	// add error status penetration
	storage.AfterGather.AppendFilter(itrace.PenetrateErrorTracing)
	// add omit certain error status list
	if len(ipt.OmitErrStatus) != 0 {
		storage.AfterGather.AppendFilter(itrace.OmitStatusCodeFilterWrapper(ipt.OmitErrStatus))
	}
	// add sampler
	if ipt.Sampler != nil {
		sampler = ipt.Sampler
	} else {
		sampler = &itrace.Sampler{SamplingRateGlobal: 1}
	}
	storage.AfterGather.AppendFilter(sampler.Sample)

	storage.GlobalTags = ipt.Tags

	if len(ipt.IgnoreAttributeKeys) > 0 {
		storage.RegexpString = strings.Join(ipt.IgnoreAttributeKeys, "|")
	}

	open := false
	// 从配置文件 开启
	if ipt.OHTTPc.Enable {
		// add option
		ipt.OHTTPc.storage = storage
		ipt.OHTTPc.ExpectedHeaders = ipt.ExpectedHeaders
		open = true
	}
	if ipt.Ogrpc.TraceEnable || ipt.Ogrpc.MetricEnable {
		open = true
		ipt.Ogrpc.ExpectedHeaders = ipt.ExpectedHeaders
		go ipt.Ogrpc.run(storage)
	}
	if open {
		// add calculators
		storage.AfterGather.AppendCalculator(itrace.StatTracingInfo)
		go storage.Run()
		for {
			select {
			case <-datakit.Exit.Wait():
				ipt.exit()
				l.Infof("%s exit", ipt.inputName)
				return

			case <-ipt.semStop.Wait():
				ipt.exit()
				l.Infof("%s return", ipt.inputName)
				return
			}
		}
	}
}

func init() { //nolint:gochecknoinits
	inputs.Add(inputName, func() inputs.Input {
		return &Input{
			inputName: inputName,
			semStop:   cliutils.NewSem(),
		}
	})
}
