package main

import (
	"github.com/erobx/trading-bot/pkg/app"
)

func main() {
	app := app.NewApp()
	app.Start()
}
