package main

import (
	"time"

	"github.com/erobx/trading-bot/pkg/sim"
)

func main() {
	quit := make(chan struct{})
	sim := sim.NewSim(quit)
	go sim.Start()
	time.Sleep(time.Second * 5)
	quit <-struct{}{}
	time.Sleep(time.Second * 1)
}
