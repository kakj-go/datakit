package etcd

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"
	"time"

	influxdb "github.com/influxdata/influxdb1-client/v2"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
)

const _PROMETHEUS_TO_METRIC_MEASUREMENT = "tmp"

var _ETCD_ACTION_LIST = map[string]byte{
	"create":       '0',
	"set":          '0',
	"get":          '0',
	"getRecursive": '0',
}

type stream struct {
	etc *Etcd
	//
	sub *Subscribe
	// http://host:port/metrics
	address string
	//
	startTime time.Time
	// mate data
	points []*influxdb.Point
}

func newStream(sub *Subscribe, etc *Etcd) *stream {
	return &stream{
		etc:       etc,
		sub:       sub,
		address:   fmt.Sprintf("http://%s:%d/metrics", sub.EtcdHost, sub.EtcdPort),
		startTime: time.Now(),
	}
}

func (s *stream) start(wg *sync.WaitGroup) error {
	defer wg.Done()

	if s.sub.Measurement == "" {
		return fmt.Errorf("invalid measurement")
	}

	ticker := time.NewTicker(time.Second * s.sub.Cycle)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.exec()
		default:
			// nil
		}
	}
}

func (s *stream) exec() error {

	resp, err := http.Get(s.address)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	metrics, err := ParseV2(resp.Body)
	if err != nil {
		return err
	}

	if len(metrics) == 0 {
		return fmt.Errorf("The metrics is nil")
	}

	var tags = make(map[string]string)
	var fields = make(map[string]interface{}, len(metrics))

	// prometheus to point
	for _, metric := range metrics {

		for k, v := range metric.Tags() {

			// 保存prometheus中4个action对应字段的fields数据，特殊需求下的产物
			// 最终结果类似：

			// prometheus,action=create etcd_debugging_store_writes_total=1 1586857769285050000
			// prometheus,action=set etcd_debugging_store_writes_total=4 1586857769285050000
			// prometheus,action=get etcd_debugging_store_reads_total=1 1586857769285050000
			// prometheus,action=getRecursive etcd_debugging_store_reads_total=2 1586857769285050000
			// ============== TO =============
			// etcd_debugging_store_writes_total_create=1
			// etcd_debugging_store_writes_total_set=4
			// etcd_debugging_store_reads_total_get=1
			// etcd_debugging_store_reads_total_getRecursive=2

			// It's STUPID!
			if _, ok := _ETCD_ACTION_LIST[v]; ok {
				for _k, _v := range metric.Fields() {
					fields[_k+"_"+v] = _v
					goto END_ACTION
				}
			}

			if _, ok := etcdCollectList[k]; ok {
				tags[k] = v
			}
		}

		for k, v := range metric.Fields() {
			if _, ok := etcdCollectList[k]; ok {
				fields[k] = v
			}
		}

	END_ACTION:
	}

	pt, err := influxdb.NewPoint(s.sub.Measurement, tags, fields, metrics[0].Time())
	if err != nil {
		return err
	}

	s.points = []*influxdb.Point{pt}

	return s.flush()
}

func (s *stream) flush() (err error) {
	// fmt.Printf("%v\n", s.points)
	err = s.etc.ProcessPts(s.points)
	s.points = nil
	return err
}

// Parse returns a slice of Metrics from a text representation of a
// metrics
func ParseV2(prodata io.Reader) ([]telegraf.Metric, error) {
	var metrics []telegraf.Metric
	var parser expfmt.TextParser

	metricFamilies, err := parser.TextToMetricFamilies(prodata)
	if err != nil {
		return nil, fmt.Errorf("reading text format failed: %s", err)
	}

	// make sure all metrics have a consistent timestamp so that metrics don't straddle two different seconds
	now := time.Now()
	// read metrics
	for metricName, mf := range metricFamilies {
		for _, m := range mf.Metric {
			// reading tags
			tags := makeLabels(m)

			if mf.GetType() == dto.MetricType_SUMMARY {
				// summary metric
				telegrafMetrics := makeQuantilesV2(m, tags, metricName, mf.GetType(), now)
				metrics = append(metrics, telegrafMetrics...)
			} else if mf.GetType() == dto.MetricType_HISTOGRAM {
				// histogram metric
				telegrafMetrics := makeBucketsV2(m, tags, metricName, mf.GetType(), now)
				metrics = append(metrics, telegrafMetrics...)
			} else {
				// standard metric
				// reading fields
				fields := make(map[string]interface{})
				fields = getNameAndValueV2(m, metricName)
				// converting to telegraf metric
				if len(fields) > 0 {
					var t time.Time
					if m.TimestampMs != nil && *m.TimestampMs > 0 {
						t = time.Unix(0, *m.TimestampMs*1000000)
					} else {
						t = now
					}
					metric, err := metric.New(_PROMETHEUS_TO_METRIC_MEASUREMENT, tags, fields, t, valueType(mf.GetType()))
					if err == nil {
						metrics = append(metrics, metric)
					}
				}
			}
		}
	}

	return metrics, err
}

// Get Quantiles for summary metric & Buckets for histogram
func makeQuantilesV2(m *dto.Metric, tags map[string]string, metricName string, metricType dto.MetricType, now time.Time) []telegraf.Metric {
	var metrics []telegraf.Metric
	fields := make(map[string]interface{})
	var t time.Time
	if m.TimestampMs != nil && *m.TimestampMs > 0 {
		t = time.Unix(0, *m.TimestampMs*1000000)
	} else {
		t = now
	}
	fields[metricName+"_count"] = float64(m.GetSummary().GetSampleCount())
	fields[metricName+"_sum"] = float64(m.GetSummary().GetSampleSum())
	met, err := metric.New(_PROMETHEUS_TO_METRIC_MEASUREMENT, tags, fields, t, valueType(metricType))
	if err == nil {
		metrics = append(metrics, met)
	}

	for _, q := range m.GetSummary().Quantile {
		newTags := tags
		fields = make(map[string]interface{})

		newTags["quantile"] = fmt.Sprint(q.GetQuantile())
		fields[metricName] = float64(q.GetValue())

		quantileMetric, err := metric.New(_PROMETHEUS_TO_METRIC_MEASUREMENT, newTags, fields, t, valueType(metricType))
		if err == nil {
			metrics = append(metrics, quantileMetric)
		}
	}
	return metrics
}

// Get Buckets  from histogram metric
func makeBucketsV2(m *dto.Metric, tags map[string]string, metricName string, metricType dto.MetricType, now time.Time) []telegraf.Metric {
	var metrics []telegraf.Metric
	fields := make(map[string]interface{})
	var t time.Time
	if m.TimestampMs != nil && *m.TimestampMs > 0 {
		t = time.Unix(0, *m.TimestampMs*1000000)
	} else {
		t = now
	}
	fields[metricName+"_count"] = float64(m.GetHistogram().GetSampleCount())
	fields[metricName+"_sum"] = float64(m.GetHistogram().GetSampleSum())

	met, err := metric.New(_PROMETHEUS_TO_METRIC_MEASUREMENT, tags, fields, t, valueType(metricType))
	if err == nil {
		metrics = append(metrics, met)
	}

	for _, b := range m.GetHistogram().Bucket {
		newTags := tags
		fields = make(map[string]interface{})
		newTags["le"] = fmt.Sprint(b.GetUpperBound())
		fields[metricName+"_bucket"] = float64(b.GetCumulativeCount())

		histogramMetric, err := metric.New(_PROMETHEUS_TO_METRIC_MEASUREMENT, newTags, fields, t, valueType(metricType))
		if err == nil {
			metrics = append(metrics, histogramMetric)
		}
	}
	return metrics
}

func valueType(mt dto.MetricType) telegraf.ValueType {
	switch mt {
	case dto.MetricType_COUNTER:
		return telegraf.Counter
	case dto.MetricType_GAUGE:
		return telegraf.Gauge
	case dto.MetricType_SUMMARY:
		return telegraf.Summary
	case dto.MetricType_HISTOGRAM:
		return telegraf.Histogram
	default:
		return telegraf.Untyped
	}
}

// Get labels from metric
func makeLabels(m *dto.Metric) map[string]string {
	result := map[string]string{}
	for _, lp := range m.Label {
		result[lp.GetName()] = lp.GetValue()
	}
	return result
}

// Get name and value from metric
func getNameAndValueV2(m *dto.Metric, metricName string) map[string]interface{} {
	fields := make(map[string]interface{})
	if m.Gauge != nil {
		if !math.IsNaN(m.GetGauge().GetValue()) {
			fields[metricName] = float64(m.GetGauge().GetValue())
		}
	} else if m.Counter != nil {
		if !math.IsNaN(m.GetCounter().GetValue()) {
			fields[metricName] = float64(m.GetCounter().GetValue())
		}
	} else if m.Untyped != nil {
		if !math.IsNaN(m.GetUntyped().GetValue()) {
			fields[metricName] = float64(m.GetUntyped().GetValue())
		}
	}
	return fields
}
