package main

import (
	"fmt"

	"hzxconf"
	"hzxdebug"
	"node"
	"ui"
	"worker"
)

var (
	conf *hzxconf.Config
)

func main() {
	debug.Init()
	conf := &debug.Conf
	
	debug.Info.Println("Running on "+debug.GetOS())
	debug.Info.Println(fmt.Sprintf("Known nodes: %s",conf.Nodes))

	debug.Info.Println(fmt.Sprintf("Starting Node on %s:%d",conf.NodeListen,conf.NodePort))
	go node.ServeNode()

	debug.Info.Println("Starting P2P processes")
	go worker.Worker()

	debug.Info.Println(fmt.Sprintf("Starting UI on %s:%d",conf.UIListen,conf.UIPort))
	ui.ServeUI(conf.UIListen,conf.UIPort)
}