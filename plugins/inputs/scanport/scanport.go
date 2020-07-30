package scanport

import (
	"bytes"
	"fmt"
	"net"
	"regexp"
	"strings"
	"sync"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

const (
	pattern = "^(?:(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}(?:[0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\/([1-9]|[1-2]\\d|3[0-1])$"
)

var (
	l    *logger.Logger
	name = "scanport"
)

func (_ *Scanport) Catalog() string {
	return "network"
}

func (_ *Scanport) SampleConfig() string {
	return configSample
}

func (_ *Scanport) Description() string {
	return ""
}

func (_ *Scanport) Gather() error {
	return nil
}

func (_ *Scanport) Init() error {
	return nil
}

func (s *Scanport) Run() {
	l = logger.SLogger("scanport")

	l.Info("scanport input started...")

	s.checkCfg()

	tick := time.NewTicker(s.IntervalDuration)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			// handle
			s.handle()
		case <-datakit.Exit.Wait():
			l.Info("exit")
			return
		}
	}
}

func (s *Scanport) checkCfg() {
	// 采集频度
	s.IntervalDuration = 10 * time.Minute

	if s.Interval != "" {
		du, err := time.ParseDuration(s.Interval)
		if err != nil {
			l.Errorf("bad interval %s: %s, use default: 10m", s.Interval, err.Error())
		} else {
			s.IntervalDuration = du
		}
	}

	// 指标集名称
	if s.Metric != "" {
		s.Metric = "scanport"
	}

	// cidr
	if s.Cidr != "" {
		matched, err := regexp.MatchString(pattern, s.Cidr)
		if err != nil {
			l.Errorf("incorrect input cidr config `%s': %s", s.Cidr, err)
		}

		if !matched {
			l.Errorf("incorrect input cidr config `%s': %s", s.Cidr, err)
		}
	}

	if len(s.Ips) == 0 {
		ipList, err := s.walkCIDR()
		if err != nil {
			l.Errorf("incorrect input cidr config `%s': %s", s.Cidr, err)
		}

		s.Ips = append(s.Ips, ipList...)
	}

	s.Timeout = time.Millisecond
}

// handle
func (s *Scanport) handle() error {
	lines := [][]byte{}

	jobCh := s.genJob()
	retCh := make(chan []byte, 200)

	wg := new(sync.WaitGroup)

	workPool(50, jobCh, retCh, wg)

	wg.Wait()
	close(retCh)
	for ret := range retCh {

		lines = append(lines, ret)
	}

	pushLines := bytes.Join(lines, []byte("\n"))
	err := io.NamedFeed(pushLines, io.Metric, name)
	if err != nil {
		l.Errorf("push metric point error %s", err)
	}
	return nil
}

func (s *Scanport) genJob() <-chan task {
	jobCh := make(chan task, 200)

	go func() {
		for _, proto := range s.Protocol {
			for _, ip := range s.Ips {
				for port := s.PortStart; port <= s.PortEnd; port++ {
					t := task{
						proto:   proto,
						metric:  s.Metric,
						ip:      ip,
						port:    port,
						timeout: s.Timeout,
					}

					jobCh <- t
				}
			}
		}
		close(jobCh)
	}()

	return jobCh
}

// cidr解析
func (s *Scanport) walkCIDR() ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(s.Cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// dial
func dial(protocol string, ip string, port int, timeout time.Duration) bool {
	conn, err := net.DialTimeout(protocol, fmt.Sprintf("%s:%d", ip, port), timeout)
	if err != nil {
		if strings.Contains(err.Error(), "too many open files") {
			l.Errorf("too many open files %v", err.Error())
		}
		return false
	}

	if conn != nil {
		defer conn.Close()
	}

	return true
}

// workpool
func workPool(n int, jobCh <-chan task, retCh chan<- []byte, wg *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		wg.Add(1)
		go work(i, jobCh, retCh, wg)
	}
}

type task struct {
	metric  string
	proto   string
	ip      string
	port    int
	timeout time.Duration
}

func work(id int, jobCh <-chan task, retCh chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobCh {
		if dial(job.proto, job.ip, job.port, job.timeout) {
			fields := make(map[string]interface{})
			tags := make(map[string]string)

			tags["protocol"] = job.proto
			fields["port"] = job.port
			fields["ip"] = job.ip

			ptline, err := io.MakeMetric(job.metric, tags, fields, time.Now())
			if err != nil {
				l.Errorf("new point failed: %s", err.Error())
			}

			retCh <- ptline
		}
	}
}

func init() {
	inputs.Add(name, func() inputs.Input {
		return &Scanport{}
	})
}
