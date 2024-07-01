package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/erobx/trading-bot/pkg/app/handler"
	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/app/service"
	"github.com/erobx/trading-bot/pkg/db"
)

type App struct {
	Mux *http.ServeMux
}

func NewApp() *App {
	return &App{
		Mux: http.NewServeMux(),
	}
}

var dev = true

func disableCacheInDevMode(next http.Handler) http.Handler {
	if !dev {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func (s *App) Start() {
	m, err := db.NewMarket()
	if err != nil {
		panic(err)
	}

	svc := service.NewMarketService(m)
	//svc = service.NewLogService(svc)

	h := handler.NewDefaultHandler(svc)

	s.Mux.Handle("/public/", disableCacheInDevMode(http.StripPrefix("/public", http.FileServer(http.Dir("public")))))
	s.Mux.Handle("/", h)

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      s.Mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	fmt.Printf("Listening on %s...\n", server.Addr)
	server.ListenAndServe()
}

func addData(m *db.Market) {
	s := model.BuildSkin()
	m.AddSkin(s)
	g := model.NewGroup("pink")
	m.AddGroup(g)
}
