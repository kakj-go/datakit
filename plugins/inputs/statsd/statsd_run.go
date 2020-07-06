package statsd

import (
	"sync"
	"bufio"
	"errors"
	"net"
	"regexp"
	"strings"
	"time"

	influxdb "github.com/influxdata/influxdb1-client/v2"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
)

var (
	ConnectionReset = errors.New("ConnectionReset")
)

func (p *StatsdParams) gather(wg *sync.WaitGroup) {
	var connectFail bool = true
	var conn net.Conn
	var err  error
	tick := time.NewTicker(time.Duration(p.input.Interval) * time.Second)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			if connectFail {
				conn, err = net.Dial("tcp", p.input.Host)
				if err != nil {
					connectFail = true
				} else {
					connectFail = false
				}
			}

			if connectFail == false && conn != nil {
				err = p.getMetrics(conn)
				if err != nil && err != ConnectionReset {
					Log.Errorf("getMetrics err: %s", err.Error())
				}

				if err == ConnectionReset {
					connectFail = true
					conn.Close()
				}
			} else {
				err = p.reportNotUp()
				if err != nil {
					Log.Errorf("reportNotUp err: %s", err.Error())
				}
			}

		case <-datakit.Exit.Wait():
			wg.Done()
			Log.Info("input statsd exit")
			return
		}
	}
}

func (p *StatsdParams) getMetrics(conn net.Conn) error {
	var pt *influxdb.Point
	var err error

	tags := make(map[string]string)
	fields := make(map[string]interface{})

	tags["host"] = p.input.Host
	fields["is_up"] = true

	err = getMetric(conn, "counters", fields)
	if err != nil {
		goto ERR
	}

	err = getMetric(conn, "gauges", fields)
	if err != nil {
		goto ERR
	}

	err = getMetric(conn, "timers", fields)
	if err != nil {
		goto ERR
	}

	fields["can_connect"] = true
	pt, err = influxdb.NewPoint(p.input.MetricName, tags, fields, time.Now())
	if err != nil {
		return err
	}
	err = p.output.IoFeed([]byte(pt.String()), io.Metric)
	if err != nil {
		return err
	}

ERR:
	fields["can_connect"] = false
	pt, _ = influxdb.NewPoint(p.input.MetricName, tags, fields, time.Now())
	err = p.output.IoFeed([]byte(pt.String()), io.Metric)

	return ConnectionReset
}

func getMetric(conn net.Conn, msg string, fields map[string]interface{}) error {
	//buf := make([]byte, 0, 1024)
	_, err := conn.Write([]byte(msg))
	if err != nil {
		return err
	}
	bio := bufio.NewReader(conn)
	s, err := bio.ReadString('}')
	if err != nil {
		return err
	}

	exp := `(?s:\{(.*)\})`
	r:= regexp.MustCompile(exp)
	matchs := r.FindStringSubmatch(s)
	if len(matchs) < 2 {
		return nil
	}

	cnt := strings.Count(matchs[1], ":")
	fields[msg+"_count"] = cnt

	return nil
}

func (p *StatsdParams) reportNotUp() error {
	tags := make(map[string]string)
	fields := make(map[string]interface{})

	tags["host"] = p.input.Host
	fields["is_up"] = false
	fields["can_connect"] = false

	pt, err := influxdb.NewPoint(p.input.MetricName, tags, fields, time.Now())
	if err != nil {
		return err
	}
	err = p.output.IoFeed([]byte(pt.String()), io.Metric)
	return err
}
