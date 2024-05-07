package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/erobx/trading-bot/pkg/app/handler"
	"github.com/erobx/trading-bot/pkg/app/service"
	"github.com/erobx/trading-bot/pkg/db"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (s *App) Start() {
	m, err := db.NewMarket()
	if err != nil {
		panic(err)
	}

	svc := service.NewMarketService(m)
	svc = service.NewLogService(svc)

	h := handler.NewDefaultHandler(svc)

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      h,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	fmt.Printf("Listening on %s...\n", server.Addr)
	server.ListenAndServe()
}
