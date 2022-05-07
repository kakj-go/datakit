// Package io implements datakits data transfer among inputs.
package io

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/cache"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/dataway"
	pb "google.golang.org/protobuf/proto"
)

const (
	minGZSize   = 1024
	maxKodoPack = 10 * 1000 * 1000

	cacheBucket = "io_upload_metric"
)

var (
	testAssert                 = false
	highFreqCleanInterval      = time.Millisecond * 500
	datawayListIntervalDefault = 50
	heartBeatIntervalDefault   = 40
	log                        = logger.DefaultSLogger("io")

	DisableLogFilter            bool
	DisableHeartbeat            bool
	DisableDatawayList          bool
	FlagDebugDisableDatawayList bool
)

type IOConfig struct {
	FeedChanSize              int                 `toml:"feed_chan_size"`
	HighFreqFeedChanSize      int                 `toml:"high_frequency_feed_chan_size"`
	MaxCacheCount             int64               `toml:"max_cache_count"`
	CacheDumpThreshold        int64               `toml:"cache_dump_threshold"`
	MaxDynamicCacheCount      int64               `toml:"max_dynamic_cache_count"`
	DynamicCacheDumpThreshold int64               `toml:"dynamic_cache_dump_threshold"`
	FlushInterval             string              `toml:"flush_interval"`
	OutputFile                string              `toml:"output_file"`
	OutputFileInputs          []string            `toml:"output_file_inputs"`
	EnableCache               bool                `toml:"enable_cache"`
	Filters                   map[string][]string `toml:"filters"`
}

type Option struct {
	CollectCost time.Duration
	HighFreq    bool
	Version     string
	HTTPHost    string
	PostTimeout time.Duration
	Sample      func(points []*Point) []*Point
}

type IO struct {
	conf *IOConfig

	dw *dataway.DataWayCfg

	in        chan *iodata
	in2       chan *iodata // high-freq chan
	inLastErr chan *lastError

	lastBodyBytes int
	SentBytes     int

	inputstats map[string]*InputsStat
	lock       sync.RWMutex

	cache        map[string][]*Point
	dynamicCache map[string][]*Point

	fd *os.File

	cacheCnt        int64
	dynamicCacheCnt int64
	droppedTotal    int64
	outputFileSize  int64
}

type IoStat struct {
	SentBytes int `json:"sent_bytes"`
}

type iodata struct {
	category, name string
	opt            *Option
	pts            []*Point
}

func TestOutput() {
	testAssert = true
}

func SetTest() {
	testAssert = true
}

//nolint:gocyclo
func (x *IO) DoFeed(pts []*Point, category, name string, opt *Option) error {
	if testAssert {
		return nil
	}

	log.Debugf("io feed %s|%s", name, category)

	ch := x.in
	if opt != nil && opt.HighFreq {
		ch = x.in2
	}

	switch category {
	case datakit.MetricDeprecated:
	case datakit.Network:
	case datakit.KeyEvent:
	case datakit.CustomObject:

	case datakit.Logging,
		datakit.Tracing,
		datakit.Metric,
		datakit.Object:
		pts = filterPts(category, pts)

	case datakit.Security:
	case datakit.RUM:
	default:
		return fmt.Errorf("invalid category `%s'", category)
	}

	// Maybe all points been filtered, but we still send the feeding into io.
	// We can still see some inputs/data are sending to io in monitor. Do not
	// optimize the feeding, or we see nothing on monitor about these filtered
	// points.
	select {
	case ch <- &iodata{
		category: category,
		pts:      pts,
		name:     name,
		opt:      opt,
	}:
	case <-datakit.Exit.Wait():
		log.Warnf("%s/%s feed skipped on global exit", category, name)
	}

	return nil
}

func (x *IO) ioStop() {
	if x.fd != nil {
		if err := x.fd.Close(); err != nil {
			log.Error(err)
		}
	}
	if x.conf.EnableCache {
		if err := cache.Stop(); err != nil {
			log.Error(err)
		}
	}
}

func (x *IO) updateStats(d *iodata) {
	now := time.Now()
	stat, ok := x.inputstats[d.name]

	if !ok {
		stat = &InputsStat{
			First: now,
		}
		x.inputstats[d.name] = stat
	}

	stat.Total += int64(len(d.pts))
	stat.Count++
	stat.Last = now
	stat.Category = d.category

	if (stat.Last.Unix() - stat.First.Unix()) > 0 {
		stat.Frequency = fmt.Sprintf("%.02f/min",
			float64(stat.Count)/(float64(stat.Last.Unix()-stat.First.Unix())/60))
	}
	stat.AvgSize = (stat.Total) / stat.Count

	if d.opt != nil {
		stat.Version = d.opt.Version
		stat.totalCost += d.opt.CollectCost
		stat.AvgCollectCost = (stat.totalCost) / time.Duration(stat.Count)
		if d.opt.CollectCost > stat.MaxCollectCost {
			stat.MaxCollectCost = d.opt.CollectCost
		}
	}
}

func (x *IO) ifMatchOutputFileInput(feedName string) bool {
	for _, v := range x.conf.OutputFileInputs {
		if v == feedName {
			return true
		}
	}
	return false
}

func (x *IO) cacheData(d *iodata, tryClean bool) {
	if d == nil {
		log.Warn("get empty data, ignored")
		return
	}

	log.Debugf("get iodata(%d points) from %s|%s", len(d.pts), d.category, d.name)

	x.updateStats(d)

	if d.opt != nil && d.opt.HTTPHost != "" {
		x.dynamicCache[d.opt.HTTPHost] = append(x.dynamicCache[d.opt.HTTPHost], d.pts...)
		x.dynamicCacheCnt += int64(len(d.pts))
	} else {
		x.cache[d.category] = append(x.cache[d.category], d.pts...)
		x.cacheCnt += int64(len(d.pts))
	}

	if x.conf.OutputFile != "" {
		bodies, err := x.buildBody(d.pts)
		if err != nil {
			log.Errorf("build iodata bodies failed: %s", err)
		}
		for _, body := range bodies {
			if len(x.conf.OutputFileInputs) == 0 || x.ifMatchOutputFileInput(d.name) {
				if err := x.fileOutput(body.buf); err != nil {
					log.Error("fileOutput: %s, ignored", err.Error())
				}
			}
		}
	}

	if (tryClean && x.conf.MaxCacheCount > 0 && x.cacheCnt > x.conf.MaxCacheCount) ||
		(x.conf.MaxDynamicCacheCount > 0 && x.dynamicCacheCnt > x.conf.MaxDynamicCacheCount) {
		x.flushAll()
	}
}

func (x *IO) cleanHighFreqIOData() {
	if len(x.in2) > 0 {
		log.Debugf("clean %d cache on high-freq-chan", len(x.in2))
	}

	for {
		select {
		case d := <-x.in2: // eat all cached data
			x.cacheData(d, false)
		default:
			return
		}
	}
}

func (x *IO) init() error {
	if x.conf.OutputFile != "" {
		f, err := os.OpenFile(x.conf.OutputFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0o644) //nolint:gosec
		if err != nil {
			log.Error(err)
			return err
		}

		x.fd = f
	}

	return nil
}

func (x *IO) StartIO(recoverable bool) {
	g := datakit.G("io")

	// start log filter
	if !DisableLogFilter {
		g.Go(func(_ context.Context) error {
			StartFilter()
			return nil
		})
	}

	g.Go(func(_ context.Context) error {
		if err := x.init(); err != nil {
			log.Errorf("init io err %v", err)
			return nil
		}

		defer x.ioStop()

		du, err := time.ParseDuration(x.conf.FlushInterval)
		if err != nil {
			l.Warnf("time.ParseDuration: %s, ignored", err)
			du = time.Second * 10
		}
		tick := time.NewTicker(du)
		defer tick.Stop()

		highFreqRecvTicker := time.NewTicker(highFreqCleanInterval)
		defer highFreqRecvTicker.Stop()

		heartBeatTick := time.NewTicker(time.Second * time.Duration(heartBeatIntervalDefault))
		defer heartBeatTick.Stop()

		datawaylistTick := time.NewTicker(time.Second * time.Duration(datawayListIntervalDefault))
		defer datawaylistTick.Stop()

		for {
			select {
			case d := <-x.in:
				x.cacheData(d, true)

			case <-highFreqRecvTicker.C:
				x.cleanHighFreqIOData()

			case <-tick.C:
				x.flushAll()
				if x.conf.EnableCache {
					x.flushCache()
					log.Debugf("cache info:%s", cache.Info())
				}

			case e := <-x.inLastErr:
				// Every inptus may report error, we append these errors into
				// input's stats info.
				x.updateLastErr(e)

			case <-heartBeatTick.C:
				log.Debugf("### enter heartBeat")
				if !DisableHeartbeat {
					heartBeatInterval, err := x.dw.HeartBeat()
					if err != nil {
						log.Warnf("dw.HeartBeat: %s, ignored", err.Error())
					}
					if heartBeatInterval != heartBeatIntervalDefault {
						heartBeatTick.Reset(time.Second * time.Duration(heartBeatInterval))
						heartBeatIntervalDefault = heartBeatInterval
					}
				}

			case <-datawaylistTick.C:
				log.Debugf("### enter dataway list")
				if !DisableDatawayList {
					var dws []string
					var err error
					var datawayListInterval int
					dws, datawayListInterval, err = x.dw.DatawayList()
					if err != nil {
						log.Warnf("DatawayList(): %s, ignored", err)
					}
					dataway.AvailableDataways = dws
					if datawayListInterval != datawayListIntervalDefault {
						datawaylistTick.Reset(time.Second * time.Duration(datawayListInterval))
						datawayListIntervalDefault = datawayListInterval
					}
				}

			case <-datakit.Exit.Wait():
				log.Info("io exit on exit")
				return nil
			}
		}
	})

	log.Info("starting...")
}

func (x *IO) updateLastErr(e *lastError) {
	x.lock.Lock()
	defer x.lock.Unlock()

	stat, ok := x.inputstats[e.from]
	if !ok {
		stat = &InputsStat{
			First: time.Now(),
			Last:  time.Now(),
		}
		x.inputstats[e.from] = stat
	}

	stat.LastErr = e.err
	stat.LastErrTS = e.ts
}

func (x *IO) flushAll() {
	x.flush()

	if x.cacheCnt > 0 {
		log.Warnf("post failed cache count: %d", x.cacheCnt)
	}

	// dump cache pts
	if x.conf.CacheDumpThreshold > 0 && x.cacheCnt > x.conf.CacheDumpThreshold {
		log.Warnf("failed cache count reach max limit(%d), cleanning cache...", x.conf.MaxCacheCount)
		for k := range x.cache {
			x.cache[k] = nil
		}
		x.droppedTotal += x.cacheCnt
		x.cacheCnt = 0
	}
	// dump dynamic cache pts
	if x.conf.DynamicCacheDumpThreshold > 0 && x.dynamicCacheCnt > x.conf.DynamicCacheDumpThreshold {
		log.Warnf("failed dynamicCache count reach max limit(%d), cleanning cache...", x.conf.MaxDynamicCacheCount)
		for k := range x.dynamicCache {
			x.dynamicCache[k] = nil
		}
		x.droppedTotal += x.dynamicCacheCnt
		x.dynamicCacheCnt = 0
	}
}

func (x *IO) flush() {
	for k, v := range x.cache {
		if err := x.doFlush(v, k); err != nil {
			log.Errorf("post %d to %s failed", len(v), k)
			if !x.conf.EnableCache {
				continue
			}

			if err := x.putCache(k, v); err != nil {
				log.Warn("failed to put cache: %s", err)
				continue
			}
		}

		if len(v) > 0 {
			x.cacheCnt -= int64(len(v))
			log.Debugf("clean %d cache on %s, remain: %d", len(v), k, x.cacheCnt)
			x.cache[k] = nil
		}
	}

	// flush dynamic cache: __not__ post to default dataway
	for k, v := range x.dynamicCache {
		if err := x.doFlush(v, k); err != nil {
			log.Errorf("post %d to %s failed", len(v), k)
			if !x.conf.EnableCache {
				// clear data
				x.dynamicCache[k] = nil
				continue
			}

			if err := x.putCache(k, v); err != nil {
				log.Warn("failed to put cache: %s", err)
				// clear data
				x.dynamicCache[k] = nil
				continue
			}
		}

		if len(v) > 0 {
			x.dynamicCacheCnt -= int64(len(v))
			log.Debugf("clean %d dynamicCache on %s, remain: %d", len(v), k, x.dynamicCacheCnt)
			x.dynamicCache[k] = nil
		}
	}
}

type body struct {
	buf  []byte
	gzon bool
}

var lines = bytes.Buffer{}

func (x *IO) buildBody(pts []*Point) ([]*body, error) {
	var (
		gz = func(lines []byte) (*body, error) {
			var (
				body = &body{buf: lines}
				err  error
			)
			log.Debugf("### io body size before GZ: %dM %dK", len(body.buf)/1000/1000, len(body.buf)/1000)
			if len(lines) > minGZSize && x.conf.OutputFile == "" {
				if body.buf, err = datakit.GZip(body.buf); err != nil {
					log.Errorf("gz: %s", err.Error())

					return nil, err
				}
				body.gzon = true
			}

			return body, nil
		}
		// lines  bytes.Buffer
		bodies []*body
	)
	lines.Reset()
	for _, pt := range pts {
		ptstr := pt.String()
		if lines.Len()+len(ptstr)+1 >= maxKodoPack {
			if body, err := gz(lines.Bytes()); err != nil {
				return nil, err
			} else {
				bodies = append(bodies, body)
			}
			lines.Reset()
		}
		lines.WriteString(ptstr)
		lines.WriteString("\n")
	}
	if body, err := gz(lines.Bytes()); err != nil {
		return nil, err
	} else {
		return append(bodies, body), nil
	}
}

func (x *IO) doFlush(pts []*Point, category string) error {
	if testAssert {
		return nil
	}

	if pts == nil {
		return nil
	}

	bodies, err := x.buildBody(pts)
	if err != nil {
		return err
	}
	for _, body := range bodies {
		if err := x.dw.Send(category, body.buf, body.gzon); err != nil {
			FeedEventLog(&DKEvent{Message: err.Error(), Status: "error", Category: "dataway"})
			return err
		}
		x.SentBytes += x.lastBodyBytes
		x.lastBodyBytes = 0
	}

	return nil
}

func (x *IO) fileOutput(body []byte) error {
	if _, err := x.fd.Write(append(body, '\n')); err != nil {
		return err
	}

	x.outputFileSize += int64(len(body))
	if x.outputFileSize > 4*1024*1024 {
		if err := x.fd.Truncate(0); err != nil {
			return err
		}
		x.outputFileSize = 0
	}

	return nil
}

func (x *IO) DroppedTotal() int64 {
	// NOTE: not thread-safe
	return x.droppedTotal
}

func (x *IO) putCache(category string, pts []*Point) error {
	bodies, err := x.buildBody(pts)
	if err != nil {
		return err
	}

	for _, body := range bodies {
		id := cliutils.XID("cache_")
		d := PBData{
			Category: category,
			Gz:       body.gzon,
			Body:     body.buf,
		}

		data, err := pb.Marshal(&d)
		if err != nil {
			return err
		}
		if err := cache.Put(cacheBucket, []byte(id), data); err != nil {
			return err
		}
		x.SentBytes += x.lastBodyBytes
		x.lastBodyBytes = 0
	}

	return nil
}

func (x *IO) flushCache() {
	log.Debugf("flush cache")
	const clean = true

	fn := func(k, v []byte) error {
		d := PBData{}
		if err := pb.Unmarshal(v, &d); err != nil {
			return err
		}

		if err := x.dw.Send(d.Category, d.Body, d.Gz); err != nil {
			FeedEventLog(&DKEvent{Message: err.Error(), Status: "error", Category: "dataway"})
			return err
		} else {
			return nil
		}
	}

	if err := cache.ForEach(cacheBucket, fn, clean); err != nil {
		log.Warnf("upload cache: %s, ignore", err)
	}
}
