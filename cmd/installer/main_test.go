package main

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	tu "gitlab.jiagouyun.com/cloudcare-tools/cliutils/testutil"
)

func TestCheckIsNewVersion(t *testing.T) {
	r := gin.New()
	r.GET("/v1/ping", func(c *gin.Context) {
		c.Data(200, "", []byte(`{ "content":{ "version": "1.2.3", "uptime": "30m", "host": "wtf" }}`))
	})

	_ = r

	ts := httptest.NewServer(r)
	time.Sleep(time.Second)
	defer ts.Close()

	cases := []struct {
		ver  string
		fail bool
	}{
		{
			ver: "1.2.3",
		},
		{
			ver:  "1.2.4",
			fail: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.ver, func(t *testing.T) {
			err := checkIsNewVersion(ts.URL, tc.ver)
			if tc.fail {
				tu.NotOk(t, err, "expect err, not nil")
				t.Logf("expect err: %s", err)
				return
			} else {
				tu.Ok(t, err)
			}
		})
	}
}

func TestCheckUpgradeVersion(t *testing.T) {
	cases := []struct {
		enableExperimental int
		id, s              string
		fail               bool
	}{
		{
			id: "normal",
			s:  "1.2.3",
		},
		{
			id: "zero-minor-version",
			s:  "1.0.3",
		},

		{
			id: "large minor version",
			s:  "1.1024.3",
		},
		{
			id:   `too-large-minor-version`,
			s:    "1.1026.3",
			fail: true,
		},
		{
			id:   `unstable-version`,
			s:    "1.3.3",
			fail: true,
		},

		{
			id:   `1.1.x-stable-rc-version`,
			s:    "1.1.9-rc1", // treat 1.1.x as stable
			fail: false,
		},

		{
			id:   `1.1.x-stable-rc-testing-version`,
			s:    "1.1.7-rc1-125-g40c4860c", // also as stable
			fail: false,
		},

		{
			id:   `1.1.x-stable-rc-hotfix-version`,
			s:    "1.1.7-rc7.1", // stable
			fail: false,
		},

		{
			id:   `invalid-version-string`,
			s:    "2.1.7.0-rc1-126-g40c4860c",
			fail: true,
		},

		{
			id:   `stable_to_unstable`,
			s:    "1.3.7",
			fail: true,
		},

		{
			id:                 `stable_to_unstable_env`,
			enableExperimental: 1,
			s:                  "1.3.7",
			fail:               false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.id, func(t *testing.T) {
			envEnableExperimental = tc.enableExperimental
			err := checkUpgradeVersion(tc.s)
			if tc.fail {
				tu.NotOk(t, err, "")
				t.Logf("expect error: %s -> %s", tc.s, err)
			} else {
				tu.Ok(t, err)
			}
		})
	}
}
