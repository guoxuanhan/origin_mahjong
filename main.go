package main

import (
	"time"

	"github.com/duanhf2012/origin/v2/node"
	_ "origin_mahjong/service/loginservice"
)

func main() {
	node.OpenProfilerReport(time.Second * 10)
	node.Start()
}
