package main

import (
	"fmt"
	"net/http"
	"time"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (s *App) start() {
	m, err := NewMarket()
	if err != nil {
		// Fail to connect to DB
		panic(err)
	}

	svc := NewMarketService(m)
	svc = NewLogService(svc)

	h := NewDefaultHandler(svc)

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      h,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	fmt.Printf("Listening on %s...\n", server.Addr)
	server.ListenAndServe()
}
