package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/kardianos/service"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/git"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/all"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/outputs/all"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/run"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/telegrafwrap"
)

var (
	flagVersion        = flag.Bool("version", false, `show verison info`)
	flagDataWay        = flag.String("dataway", ``, `dataway IP:Port`)
	flagCheckConfigDir = flag.Bool("check-config-dir", false, `check datakit conf.d, list configired and mis-configured collectors`)
	flagInputFilters   = flag.String("input-filter", "", "filter the inputs to enable, separator is :")
	flagListCollectors = flag.Bool("tree", false, `list vailable collectors`)
)

var (
	stopCh     chan struct{} = make(chan struct{})
	waitExitCh chan struct{} = make(chan struct{})

	inputFilters = []string{}
	l            *logger.Logger
)

func main() {

	flag.Parse()

	applyFlags()

	loadConfig()

	svcConfig := &service.Config{
		Name: datakit.ServiceName,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		l.Fatal(err)
		return
	}

	l.Info("starting datakit service")

	if err = s.Run(); err != nil {
		l.Fatal(err)
	}
}

func applyFlags() {

	if *flagVersion {
		fmt.Printf(`
       Version: %s
        Commit: %s
        Branch: %s
 Build At(UTC): %s
Golang Version: %s
      Uploader: %s
`, git.Version, git.Commit, git.Branch, git.BuildAt, git.Golang, git.Uploader)
		os.Exit(0)
	}

	if *flagListCollectors {
		collectors := map[string][]string{}

		for k, v := range inputs.Inputs {
			cat := v().Catalog()
			collectors[cat] = append(collectors[cat], k)
		}

		ndatakit := 0
		for k, vs := range collectors {
			fmt.Println(k)
			for _, v := range vs {
				fmt.Printf("  |--[d] %s\n", v)
				ndatakit++
			}
		}

		nagent := 0
		collectors = map[string][]string{}
		for k, v := range config.SupportsTelegrafMetricNames {
			collectors[v.Catalog] = append(collectors[v.Catalog], k)
		}

		for k, vs := range collectors {
			fmt.Println(k)
			for _, v := range vs {
				fmt.Printf("  |--[t] %s\n", v)
				nagent++
			}
		}

		fmt.Println("===================================")
		fmt.Printf("total: %d, datakit: %d, agent: %d\n", ndatakit+nagent, ndatakit, nagent)

		os.Exit(0)
	}

	if *flagCheckConfigDir {
		config.CheckConfd()
		os.Exit(0)
	}

	if *flagInputFilters != "" {
		inputFilters = strings.Split(":"+strings.TrimSpace(*flagInputFilters)+":", ":")
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run(s)
	return nil
}

func (p *program) run(s service.Service) {
	__run()
}

func (p *program) Stop(s service.Service) error {
	close(stopCh)

	// We must wait here:
	// On windows, we stop datakit in services.msc, if datakit process do not
	// echo to here, services.msc will complain the datakit process has been
	// exit unexpected
	<-waitExitCh

	return nil
}

func exitDatakit() {
	datakit.Exit.Close()

	l.Info("wait all goroutines exit...")
	datakit.WG.Wait()

	l.Info("closing waitExitCh...")
	close(waitExitCh)
}

func __run() {

	datakit.WG.Add(1)
	go func() {
		defer datakit.WG.Done()
		if err := runTelegraf(); err != nil {
			l.Fatalf("fail to start sub service: %s", err)
		}

		l.Info("telegraf process exit ok")
	}()

	l.Info("datakit start...")
	if err := runDatakit(); err != nil && err != context.Canceled {
		l.Fatalf("datakit abort: %s", err)
	}

	l.Info("datakit start ok. Wait signal or service stop...")

	// NOTE:
	// Actually, the datakit process been managed by system service, no matter on
	// windows/UNIX, datakit should exit via `service-stop' operation, so the signal
	// branch should not reached, but for daily debugging(ctrl-c), we kept the signal
	// exit option.
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-signals:
		if sig == syscall.SIGHUP {
			// TODO: reload configures
		} else {
			l.Infof("get signal %v, wait & exit", sig)
			exitDatakit()
		}
	case <-stopCh:
		l.Infof("service stopping")
		exitDatakit()
	}

	l.Info("datakit exit.")
}

func loadConfig() {

	config.Cfg.InputFilters = inputFilters

	if err := config.LoadCfg(); err != nil {
		panic(fmt.Sprintf("load config failed: %s", err))
	}

	l = logger.SLogger("main")
	//l.Infof("input fileters %v", inputFilters)
}

func runTelegraf() error {
	telegrafwrap.Svr.Cfg = config.Cfg
	return telegrafwrap.Svr.Start()
}

func runDatakit() error {

	ag, err := run.NewAgent()
	if err != nil {
		return err
	}

	return ag.Run()
}
