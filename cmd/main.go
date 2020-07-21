package main

import (
	"csi-demo-driver/pkg/demo"
	"flag"
	"k8s.io/klog"
)

var (
	endpoint string
	nodeID   string
)

func main() {
	flag.StringVar(&endpoint, "endpoint", "", "CSI Endpoint")
	flag.StringVar(&nodeID, "nodeid", "", "node id")

	klog.InitFlags(nil)
	flag.Parse()

	d := demo.NewDriver(nodeID, endpoint)
	d.Run()
}
