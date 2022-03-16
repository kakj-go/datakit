package worker

import (
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/funcs"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/parser"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/pipeline/scriptstore"
)

type tagfield struct {
	measurement string
	dropped     bool
	tags        map[string]string
	fields      map[string]interface{}
	ts          time.Time
}

type taskData struct {
	tags        map[string]string
	fields      map[string]interface{}
	content     []string
	contentByte [][]byte
	callback    func(r []*io.Point) error
	encode      string
	dataType    string
}

func (d *taskData) GetContentStr() []string {
	return d.content
}

func (d *taskData) GetContentByte() [][]byte {
	return d.contentByte
}

func (d *taskData) ContentType() string {
	return d.dataType
}

func (d *taskData) ContentEncode() string {
	return d.encode
}

func (d *taskData) Callback(task *Task, result []*pipeline.Result) error {
	pts := []*io.Point{}
	result = ResultUtilsLoggingProcessor(task, result, d.tags, d.fields)
	ts := task.TS
	if ts.IsZero() {
		ts = time.Now()
	}
	result = ResultUtilsAutoFillTime(result, ts)
	for _, r := range result {
		if pt, err := r.MakePointIgnoreDropped(task.Source, 0, ""); err != nil {
			l.Error(err)
		} else if pt != nil {
			pts = append(pts, pt)
		}
	}
	return d.callback(pts)
}

func TestRunAsTask(t *testing.T) {
	cases := []struct {
		title       string
		measurement string
		service     string
		category    string
		plScript    string
		content     []string
	}{
		{
			title:       "case 1",
			measurement: "aa",
			category:    datakit.Logging,
			plScript: `
			json(_, time)
			set_tag(bb, "aa0")
			default_time(time)
			json(_, source)
			set_tag(source)
			`,
			content: []string{
				`{"time":"02/Dec/2021:11:55:34 +0800", "source": "bb"}`,
				`{"time":"02/Dec/2021:11:55:35 +0800"}`,
				`{"time":"02/Dec/2021:11:55:36 +0800"}`,
			},
		},
	}

	result := [][]tagfield{
		{
			{
				measurement: "bb",
				dropped:     false,
				tags: map[string]string{
					"bb":      "aa0",
					"service": "aa",
				},
				fields: map[string]interface{}{
					"message": `{"time":"02/Dec/2021:11:55:34 +0800", "source": "bb"}`,
					"status":  "info",
				},
				ts: time.Unix(1638417334, 0),
			},
			{
				measurement: "aa",
				dropped:     false,
				tags: map[string]string{
					"bb":      "aa0",
					"service": "aa",
				},
				fields: map[string]interface{}{
					"message": `{"time":"02/Dec/2021:11:55:35 +0800"}`,
					"status":  "info",
				},
				ts: time.Unix(1638417335, 0),
			},
			{
				measurement: "aa",
				dropped:     false,
				tags: map[string]string{
					"service": "aa",
					"bb":      "aa0",
				},
				fields: map[string]interface{}{
					"message": `{"time":"02/Dec/2021:11:55:36 +0800"}`,
					"status":  "info",
				},
				ts: time.Unix(1638417336, 0),
			},
		},
	}

	for index, data := range cases {
		ng, err := ParsePlScript(data.plScript)
		if err != nil && result[index] != nil {
			t.Error(data.title, ": ", err.Error())
		}

		switch data.category {
		case datakit.Logging:
			if data.service == "" {
				data.service = data.measurement
			}
		default:
		}

		r := RunAsPlTask(data.category, data.measurement, data.service, ContentString, data.content, nil, "", ng)
		if len(result[index]) != len(r) {
			t.Error("length not equal")
		}
		for k, v := range r {
			// tags fields
			assert.Equal(t, result[index][k].tags, v.GetTags())
			assert.Equal(t, result[index][k].fields, v.GetFields())

			// dropped flag
			assert.Equal(t, result[index][k].dropped, v.IsDropped())

			// measurement
			assert.Equal(t, result[index][k].measurement, v.GetMeasurement())

			// ts
			getTS, _ := v.GetTime()
			if !getTS.Equal(result[index][k].ts) {
				t.Error("time not equal")
			}
		}
	}
}

func TestNewEmptyNg(t *testing.T) {
	ng, err := parser.NewEngine("if true{}", funcs.FuncsMap, funcs.FuncsCheckMap, true)
	if err != nil {
		t.Error(err)
		return
	}
	in := "aaa"
	_ = ng.Run(in)
	v, _ := ng.Data.GetContent("message")
	if v != in {
		t.Error(v)
	}
}

func TestAddStatus(t *testing.T) {
	v := &pipeline.Result{
		Output: &parser.Output{
			Tags: map[string]string{},
			Fields: map[string]interface{}{
				"status": "WARN",
			},
		},
	}
	PPAddSatus(v, false)
	assert.Equal(t, "warning", v.Output.Fields["status"])

	v = &pipeline.Result{
		Output: &parser.Output{
			Tags: map[string]string{},
			Fields: map[string]interface{}{
				"status": "ERROR",
			},
		},
	}
	PPAddSatus(v, true)
	assert.Equal(t, v.Output.Fields, map[string]interface{}{"status": "ERROR"})
}

func TestIgnoreStatus(t *testing.T) {
	if !PPIgnoreStatus("info", []string{"info", "waring", "error"}) {
		t.Error("info")
	}
}

func TestWorker(t *testing.T) {
	ts := time.Now()
	ptCh := make(chan []*io.Point)
	// set feed func for test
	feedResult := func(pts []*io.Point) error {
		ptCh <- pts
		return nil
	}
	getResult := func() []*io.Point {
		return <-ptCh
	}

	checkUpdateDebug = time.Second
	// init manager
	InitManager(1)
	_ = os.WriteFile("/tmp/nginx-time.p", []byte(`
	json(_, time)
	set_tag(bb, "aa0")
	default_time(time)
	`), os.FileMode(0o755))
	scriptstore.LoadDotPScript2StoreWithNS(scriptstore.DefaultScriptNS,
		[]string{"/tmp/nginx-time.p"}, "")
	_ = os.Remove("/tmp/nginx-time.p")

	// 测试 panic 触发
	FeedPipelineTask(&Task{})

	cases := []*Task{
		{
			TaskName: "nginx-test-log1",
			Source:   "nginx123",

			Opt: &TaskOpt{IgnoreStatus: []string{"warn"}},
			Data: &taskData{
				dataType: ContentString,
				tags: map[string]string{
					"tk": "aaa",
				},
				content:  []string{`{"time":"02/Dec/2021:11:55:34 +0800"}`},
				callback: feedResult,
			},

			TS: ts,
		},
		{
			ScriptName: "nginx-time.p",
			TaskName:   "nginx-test-log2",
			Source:     "nginx-time",
			Data: &taskData{
				dataType: ContentString,
				tags: map[string]string{
					"tk": "aaa",
					"bb": "aa0",
				},
				content: []string{
					`{"time":"02/Dec/2021:11:55:34 +0800"}`,
					`{"time":"02/Dec/2021:11:55:35 +0800"}`,
				},
				callback: feedResult,
			},
			TS: ts,
		},
		{ // index == 2， 变更脚本
			TaskName: "nginx-test-log3",
			Source:   "nginx-time",
			Data: &taskData{
				dataType: ContentString,
				tags: map[string]string{
					"tk": "aaa",
				},
				content: []string{
					`{"time":"02/Dec/2021:11:55:34 +0800", "status":"DEBUG"}`,
					`{"time":"02/Dec/2021:11:55:35 +0800", "status":"DEBUG"}`,
				},
				callback: feedResult,
			},
			TS: ts,
		},
		{
			TaskName: "nginx-test-log4",
			Source:   "nginx-time",
			Data: &taskData{
				dataType: ContentString,
				tags: map[string]string{
					"tk": "aaa",
				},
				content: []string{
					`{"time":"02/Dec/2021:11:55:11 +0800", "status":"DEBUG"}`,
				},
				callback: feedResult,
			},

			Opt: &TaskOpt{
				IgnoreStatus: []string{"debug"},
			},
			TS: ts,
		},

		// time sub
		{
			TaskName: "time sub",
			Source:   "xxxxxx",
			Data: &taskData{
				dataType: ContentString,
				tags: map[string]string{
					"tk": "aaa",
				},
				content: []string{
					`{"timex":"02/Dec/2021:11:55:34 +0800"}`,
					`{"timex":"02/Dec/2021:11:55:35 +0800"}`,
				},
				callback: feedResult,
			},
			TS: ts,
		},
	}

	expected := []([]tagfield){
		[]tagfield{
			{
				tags: map[string]string{
					"tk": "aaa",
				},
				fields: map[string]interface{}{
					"message": `{"time":"02/Dec/2021:11:55:34 +0800"}`,
					"status":  "info",
				},
				ts: ts.Add(-time.Nanosecond),
			},
		},
		[]tagfield{
			{
				tags: map[string]string{
					"tk": "aaa",
					"bb": "aa0",
				},
				fields: map[string]interface{}{
					"message": `{"time":"02/Dec/2021:11:55:34 +0800"}`,
					"status":  "info",
				},
				ts: time.Unix(1638417334, 0),
			},
			{
				tags: map[string]string{
					"bb": "aa0",
					"tk": "aaa",
				},
				fields: map[string]interface{}{
					"message": `{"time":"02/Dec/2021:11:55:35 +0800"}`,
					"status":  "info",
				},
				ts: time.Unix(1638417335, 0),
			},
		},
		[]tagfield{
			{
				tags: map[string]string{
					"tk": "aaa",
				},
				fields: map[string]interface{}{
					"message": `{"time":"02/Dec/2021:11:55:34 +0800", "status":"DEBUG"}`,
					"status":  "debug",
				},
				ts: time.Unix(1638417334, 0),
			},
			{
				tags: map[string]string{
					"tk": "aaa",
				},
				fields: map[string]interface{}{
					"message": `{"time":"02/Dec/2021:11:55:35 +0800", "status":"DEBUG"}`,
					"status":  "debug",
				},
				ts: time.Unix(1638417335, 0),
			},
		},
		[]tagfield{},
		[]tagfield{
			{
				tags: map[string]string{
					"tk": "aaa",
				},
				fields: map[string]interface{}{
					"message": `{"timex":"02/Dec/2021:11:55:34 +0800"}`,
					"status":  "info",
				},
				ts: ts.Add(time.Nanosecond * -2),
			},
			{
				tags: map[string]string{
					"tk": "aaa",
				},
				fields: map[string]interface{}{
					"message": `{"timex":"02/Dec/2021:11:55:35 +0800"}`,
					"status":  "info",
				},
				ts: ts.Add(time.Nanosecond * -1),
			},
		},
	}

	for k, v := range cases {
		if k == 2 {
			_ = scriptstore.AppendScript(scriptstore.GitRepoScriptNS, "nginx-time.p", `
			json(_, time)
			json(_, status)
			default_time(time)`, true)
			time.Sleep(time.Second + time.Millisecond*10)
		}
		_ = FeedPipelineTask(v)
		pts := getResult()
		expectedItem := expected[k]
		t.Logf("case %d", k)
		t.Log(expectedItem)
		t.Log(pts)
		if len(pts) != len(expectedItem) {
			t.Error("count not equal")
			continue
		}
		for k2, v2 := range expectedItem {
			assert.Equal(t, v2.tags, pts[k2].Tags())
			f, _ := pts[k2].Fields()
			assert.Equal(t, v2.fields, f)
			assert.Equal(t, v2.ts.UnixNano(), pts[k2].Time().UnixNano(),
				fmt.Sprintf("index: %d %d", k, k2))
		}
	}

	datakit.Exit.Close()
	err := FeedPipelineTask(&Task{})
	time.Sleep(time.Millisecond * 100)
	if !(errors.Is(err, ErrTaskChClosed) || err == nil) {
		t.Error(err)
	}
}

func TestGrokStack(t *testing.T) {
	pl := `
	add_pattern("aa", "\\d{2}")
	grok(_, "%{aa:aa}")
	if false {
	
	} else {
		add_pattern("bb", "[a-z]{3}")
		if aa == "11" {
			add_pattern("cc", "end1")
			grok(_, "%{aa:aa},%{bb:bb},%{cc:cc}")
		} elif aa == "22" {
			grok(_, "%{aa:aa},%{bb:bb},%{INT:cc}")
		} elif aa == "33" {
			add_pattern("bb", "[\\d]{5}")	# 此处覆盖 bb 失败
			add_pattern("cc", "end3")
			grok(_, "%{aa:aa},%{bb:bb},%{cc:cc}")	
		}
	}
	`

	cases := []string{
		"11,abc,end1",
		"22,abc,end1",
		"33,abc,end3",
	}

	result := []parser.Output{
		{
			Fields: map[string]interface{}{
				"aa":      "11",
				"bb":      "abc",
				"cc":      "end1",
				"message": "11,abc,end1",
			},
		},
		{
			Fields: map[string]interface{}{
				"aa":      "22",
				"message": "22,abc,end1",
			},
		},
		{
			Fields: map[string]interface{}{
				"aa":      "33",
				"bb":      "abc",
				"cc":      "end3",
				"message": "33,abc,end3",
			},
		},
	}

	ng, err := parser.NewEngine(pl, funcs.FuncsMap, funcs.FuncsCheckMap, false)
	if err != nil {
		t.Error(err)
		return
	}
	for k, v := range cases {
		ng.Run(v)
		ret := ng.Result()
		t.Log("case", k)
		assert.Equal(t, result[k].Fields, ret.Fields)
	}
}

func BenchmarkPpWorker_Run(b *testing.B) {
	// init manager
	InitManager(-1)

	ts := time.Now()

	for i := 0; i < b.N; i++ {
		err := FeedPipelineTaskBlock(&Task{
			TaskName: "nginx-test-log",
			Source:   "nginx",
			Opt:      &TaskOpt{IgnoreStatus: []string{"warn"}},
			Data: &taskData{
				tags: map[string]string{
					"tk": "aaa",
				},
				content: []string{
					`127.0.0.1 - - [16/Dec/2021:17:25:29 +0800] "GET / HTTP/1.1" 404 162 "-" "Wget/1.20.3 (linux-gnu)"`,
				},
				callback: func(r []*io.Point) error { return nil },
			},

			TS: time.Now(),
		})
		if err != nil {
			b.Log(err)
		}
	}
	if len(taskCh) != 0 {
		time.Sleep(time.Millisecond * 100)
	}

	if len(taskCh) == 0 {
		b.Log(time.Since(ts))
	}
}
