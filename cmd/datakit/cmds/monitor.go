package cmds

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/term"

	markdown "github.com/MichaelMure/go-term-markdown"
	dkhttp "gitlab.jiagouyun.com/cloudcare-tools/datakit/http"
)

func cmdMonitor(interval time.Duration, addrStr string, verbose bool) {
	addr := "http://localhost:9529/stats"
	if addrStr != "" {
		addr = "http://" + addrStr + "/stats"
	}

	if interval < time.Second {
		interval = time.Second
	}

	run := func() {
		fmt.Print("\033[H\033[2J") // clean screen

		x, err := doCMDMonitor(addr, verbose)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(string(x))
			fmt.Printf("(Refresh at %s)Press ctrl+c to exit.\n", interval)
		}
	}

	run() // run before sleep

	tick := time.NewTicker(interval)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			run()
		}
	}
}

func doCMDMonitor(url string, verbose bool) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("%s", string(body))
	}

	ds := dkhttp.DatakitStats{
		DisableMonofont: true,
	}
	if err := json.Unmarshal(body, &ds); err != nil {
		return nil, err
	}

	l.Debugf("stats.ReloadInfo: %s", ds.ReloadInfo)

	mdtxt, err := ds.Markdown("", verbose)
	if err != nil {
		return nil, err
	}

	width := 100
	if term.IsTerminal(0) {
		w, _, err := term.GetSize(0)
		if err == nil {
			width = w
		}
	}

	if err != nil {
		return nil, err
	} else {
		if len(mdtxt) == 0 {
			return nil, fmt.Errorf("no monitor info available")
		} else {
			result := markdown.Render(string(mdtxt), width, 2)
			return result, nil
		}
	}
}
