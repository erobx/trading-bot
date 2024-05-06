package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"time"

	"github.com/erobx/trading-bot/pkg/sim"
)

func main() {
	for i := 0; i < 10; i++ {
		determineChange()
	}
}

func runSim() {
	quit := make(chan struct{})
	sim := sim.NewSim(quit)
	go sim.Start()
	time.Sleep(time.Second * 5)
	quit <- struct{}{}
	time.Sleep(time.Second * 1)
}

func determineChange() {
	hist := []float64{10.12, 9.69, 12.32}
	t := 12.12

	for i := 0; i < 100; i++ {
		hist = append(hist, randomSale())
		sort.Slice(hist, func(i, j int) bool {
			return hist[i] < hist[j]
		})
		change := (hist[len(hist)/2] - t) / t
		x := t * change
		t = t + x
	}
	fmt.Printf("Price: $%.2f\n", t)
}

func randomSale() float64 {
	min_d := float64(9.12)
	max_d := float64(12.32)
	return min_d + rand.Float64()*(max_d-min_d)
}
