// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at Guance Cloud (https://www.guance.com/).
// Copyright 2021-present Guance, Inc.

// Package sqlserver collects SQL Server metrics.
package sqlserver

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"strings"
	"time"

	mssql "github.com/denisenkom/go-mssqldb"
	"github.com/denisenkom/go-mssqldb/msdsn"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils"
	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/internal/tailer"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io/point"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
)

var _ inputs.ElectionInput = (*Input)(nil)

func (n *Input) ElectionEnabled() bool {
	return n.Election
}

func (n *Input) Pause() error {
	tick := time.NewTicker(inputs.ElectionPauseTimeout)
	defer tick.Stop()
	select {
	case n.pauseCh <- true:
		return nil
	case <-tick.C:
		return fmt.Errorf("pause %s failed", inputName)
	}
}

func (n *Input) Resume() error {
	tick := time.NewTicker(inputs.ElectionResumeTimeout)
	defer tick.Stop()
	select {
	case n.pauseCh <- false:
		return nil
	case <-tick.C:
		return fmt.Errorf("resume %s failed", inputName)
	}
}

func (*Input) SampleConfig() string {
	return sample
}

func (*Input) Catalog() string {
	return catalogName
}

func (*Input) AvailableArchs() []string {
	return datakit.AllOS
}

func (*Input) PipelineConfig() map[string]string {
	pipelineMap := map[string]string{
		inputName: pipeline,
	}
	return pipelineMap
}

//nolint:lll
func (n *Input) LogExamples() map[string]map[string]string {
	return map[string]map[string]string{
		inputName: {
			"SQLServer log": `2021-05-28 10:46:07.78 spid10s     0 transactions rolled back in database 'msdb' (4:0). This is an informational message only. No user action is required`,
		},
	}
}

func (n *Input) GetPipeline() []*tailer.Option {
	return []*tailer.Option{
		{
			Source:  inputName,
			Service: inputName,
			Pipeline: func() string {
				if n.Log != nil {
					return n.Log.Pipeline
				}
				return ""
			}(),
		},
	}
}

func (n *Input) initDB() error {
	connStr := fmt.Sprintf("sqlserver://%s:%s@%s?dial+timeout=3", n.User, n.Password, n.Host)
	cfg, _, err := msdsn.Parse(connStr)
	if err != nil {
		return err
	}
	if n.AllowTLS10 {
		// Because go1.18 defaults client-sids's TLS minimum version to TLS 1.2,
		// we need to configure MinVersion manually to enable TLS 1.0 and TLS 1.1.
		cfg.TLSConfig.MinVersion = tls.VersionTLS10
	}
	conn := mssql.NewConnectorConfig(cfg)
	db := sql.OpenDB(conn)

	if err := db.Ping(); err != nil {
		return err
	}

	n.db = db
	return nil
}

func (n *Input) RunPipeline() {
	if n.Log == nil || len(n.Log.Files) == 0 {
		return
	}

	if n.Log.Pipeline == "" {
		n.Log.Pipeline = "sqlserver.p" // use default
	}

	opt := &tailer.Option{
		Source:            inputName,
		Service:           inputName,
		Pipeline:          n.Log.Pipeline,
		GlobalTags:        n.Tags,
		IgnoreStatus:      n.Log.IgnoreStatus,
		CharacterEncoding: n.Log.CharacterEncoding,
		MultilinePatterns: []string{`^\d{4}-\d{2}-\d{2}`},
	}

	var err error
	n.tail, err = tailer.NewTailer(n.Log.Files, opt)
	if err != nil {
		l.Error(err)
		io.FeedLastError(inputName, err.Error())
		return
	}

	go n.tail.Start()
}

func (n *Input) Run() {
	l = logger.SLogger(inputName)
	l.Info("sqlserver start")

	n.Interval.Duration = config.ProtectedInterval(minInterval, maxInterval, n.Interval.Duration)

	tick := time.NewTicker(n.Interval.Duration)
	defer tick.Stop()

	n.initDBFilterMap()

	// Init DB until OK.
	for {
		if err := n.initDB(); err != nil {
			l.Errorf("initDB: %s", err.Error())
			io.FeedLastError(inputName, err.Error())
		} else {
			break
		}

		select {
		case <-tick.C:
		case <-datakit.Exit.Wait():
			l.Info("sqlserver exit")
			return
		case n.pause = <-n.pauseCh:
			// nil
		}
	}

	defer func() {
		if err := n.db.Close(); err != nil {
			l.Warnf("Close: %s", err)
		}

		if n.tail != nil {
			n.tail.Close()
		}
	}()

	for {
		if n.pause {
			l.Debugf("not leader, skipped")
		} else {
			n.getMetric()
			if len(collectCache) > 0 {
				err := io.Feed(inputName, datakit.Metric, collectCache, &io.Option{CollectCost: time.Since(n.start)})
				collectCache = collectCache[:0]
				if err != nil {
					n.lastErr = err
					l.Errorf(err.Error())
					continue
				}
			}

			if n.lastErr != nil {
				io.FeedLastError(inputName, n.lastErr.Error())
				n.lastErr = nil
			}

			select {
			case <-tick.C:
			case <-datakit.Exit.Wait():
				l.Info("sqlserver exit")
				return

			case <-n.semStop.Wait():
				n.exit()
				l.Info("sqlserver return")
				return
			}
		}
	}
}

func (n *Input) exit() {
	if n.tail != nil {
		n.tail.Close()
		l.Info("sqlserver log exit")
	}
}

func (n *Input) Terminate() {
	if n.semStop != nil {
		n.semStop.Close()
	}
}

func (n *Input) getMetric() {
	start := time.Now()
	n.start = start
	for _, v := range query {
		n.handRow(v, start)
	}
}

func (n *Input) handRow(query string, ts time.Time) {
	rows, err := n.db.Query(query)
	if err != nil {
		l.Error(err.Error())
		n.lastErr = err
		return
	}
	defer rows.Close() //nolint:errcheck

	if err := rows.Err(); err != nil {
		l.Errorf("rows.Err: %s", err)
		return
	}

	OrderedColumns, err := rows.Columns()
	if err != nil {
		l.Error(err.Error())
		n.lastErr = err
		return
	}

	for rows.Next() {
		var columnVars []interface{}
		// var fields = make(map[string]interface{})
		// store the column name with its *interface{}
		columnMap := make(map[string]*interface{})

		for _, column := range OrderedColumns {
			columnMap[column] = new(interface{})
		}
		// populate the array of interface{} with the pointers in the right order
		for i := 0; i < len(columnMap); i++ {
			columnVars = append(columnVars, columnMap[OrderedColumns[i]])
		}
		// deconstruct array of variables and send to Scan
		err := rows.Scan(columnVars...)
		if err != nil {
			l.Error(err.Error())
			n.lastErr = err
			return
		}
		measurement := ""
		tags := make(map[string]string)
		for k, v := range n.Tags {
			tags[k] = v
		}
		fields := make(map[string]interface{})
		for header, val := range columnMap {
			if str, ok := (*val).(string); ok {
				if header == "measurement" {
					measurement = str
					continue
				}
				tags[header] = strings.TrimSuffix(str, "\\")
			} else {
				if *val == nil {
					continue
				}
				fields[header] = *val
			}
		}
		if len(fields) == 0 {
			continue
		}
		if n.filterOutDBName(tags) {
			continue
		}

		point, err := point.NewPoint(measurement, tags, fields, point.MOptElectionV2(n.Election))
		if err != nil {
			l.Errorf("make point err:%s", err.Error())
			n.lastErr = err
			continue
		}
		collectCache = append(collectCache, point)
	}
}

// filterOutDBName filters out metrics according to their database_name tag.
// Metrics with database_name tag specified in db_filter are filtered out and not fed to IO.
func (n *Input) filterOutDBName(tags map[string]string) bool {
	if len(n.dbFilterMap) == 0 {
		return false
	}
	db, has := tags["database_name"]
	if !has {
		return false
	}
	if _, filterOut := n.dbFilterMap[db]; filterOut {
		l.Debugf("filter out metric from db: %s", db)
		return true
	}
	return false
}

func (n *Input) initDBFilterMap() {
	if n.dbFilterMap == nil {
		n.dbFilterMap = make(map[string]struct{}, len(n.DBFilter))
	}
	for _, db := range n.DBFilter {
		n.dbFilterMap[db] = struct{}{}
	}
}

func (n *Input) SampleMeasurement() []inputs.Measurement {
	return []inputs.Measurement{
		&ServerProperties{},
		&Performance{},
		&WaitStatsCategorized{},
		&DatabaseIO{},
		&Schedulers{},
	}
}

func init() { //nolint:gochecknoinits
	inputs.Add(inputName, func() inputs.Input {
		s := &Input{
			Interval:    datakit.Duration{Duration: time.Second * 10},
			Election:    true,
			pauseCh:     make(chan bool, inputs.ElectionPauseChannelLength),
			semStop:     cliutils.NewSem(),
			dbFilterMap: make(map[string]struct{}, 0),
		}
		return s
	})
}
