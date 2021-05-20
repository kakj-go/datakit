package cmds

import (
	"fmt"
	nhttp "net/http"

	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/cmd/installer/install"
)

const (
	dataUrl = "https://zhuyun-static-files-production.oss-cn-hangzhou.aliyuncs.com/datakit/data.tar.gz"
)

func UpdateIpDB(port int, addr string) error {
	if addr == "" {
		addr = dataUrl
	}

	fmt.Printf("Start downloading data.tar.gz...\n")

	if err := install.Download(addr, datakit.InstallDir, true, false); err != nil {
		return err
	}

	fmt.Printf("Download and decompress data.tar.gz successfully.\n")

	_, err := nhttp.Get(fmt.Sprintf("http://127.0.0.1:%d/reload", port))
	if err != nil {
		fmt.Printf("Reload datakit fail!\n")
		fmt.Printf("You need restart datakit by `datakit --restart` to make effect.\n")
		return nil
	} else {
		fmt.Printf("Update successfully.\n")
	}

	return nil
}