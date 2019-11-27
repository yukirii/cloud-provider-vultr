package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/yukirii/vultr-cloud-provider/vultr"
	"k8s.io/apiserver/pkg/server/healthz"
	"k8s.io/component-base/logs"
	"k8s.io/klog"
	"k8s.io/kubernetes/cmd/cloud-controller-manager/app"
	//_ "k8s.io/kubernetes/pkg/client/metrics/prometheus" // for client metric registration
	//_ "k8s.io/kubernetes/pkg/version/prometheus"        // for version metric registration
)

func init() {
	healthz.InstallHandler(http.DefaultServeMux)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	command := app.NewCloudControllerManagerCommand()

	logs.InitLogs()
	defer logs.FlushLogs()

	klog.V(1).Infof("vultr-cloud-controller-manager")

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
