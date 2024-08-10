package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/erobx/trading-bot/pkg/app/handler"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	d, err := db.NewConn()
	if err != nil {
		panic(err)
	}

	h := handler.NewDefaultHandler(d)
	th := handler.NewTradeupsHandler(d, time.Now)
	mh := handler.NewModalHandler(d)
	sh := handler.NewSignupHandler(d)
	lh := handler.NewLoginHandler()
	ad := handler.NewAdminHandler(d)

	username := os.Getenv("AUTH_USERNAME")
	password := os.Getenv("AUTH_PASSWORD")
	ah := NewAuth(ad, username, password)

	s.Mux.Handle("/public/", disableCacheInDevMode(http.StripPrefix("/public", http.FileServer(http.Dir("public")))))

	s.Mux.Handle("/", h)
	s.Mux.Handle("/tradeups", th)
	s.Mux.Handle("/modal", mh)
	s.Mux.Handle("/signup", sh)
	s.Mux.Handle("/login", lh)
	s.Mux.Handle("/admin", ah)

	server := &http.Server{
		Addr:         "localhost:3000",
		Handler:      s.Mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}
	fmt.Printf("Listening on %s...\n", server.Addr)
	server.ListenAndServe()
}
