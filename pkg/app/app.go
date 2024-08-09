package app

import (
	"crypto/sha256"
	"crypto/subtle"
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

type Auth struct {
	next     http.Handler
	username string
	password string
}

func NewAuth(n http.Handler, username, password string) http.Handler {
	auth := Auth{
		next:     n,
		username: username,
		password: password,
	}
	return auth
}

func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsername := sha256.Sum256([]byte(a.username))
		expectedPasword := sha256.Sum256([]byte(a.password))

		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsername[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasword[:]) == 1)

		if usernameMatch && passwordMatch {
			a.next.ServeHTTP(w, r)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}

func (s *App) Start() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	m, err := db.NewMarketConn()
	if err != nil {
		panic(err)
	}

	h := handler.NewDefaultHandler(m)
	th := handler.NewTradeupsHandler(m, time.Now)
	mh := handler.NewModalHandler(m)
	sh := handler.NewSignupHandler(m)
	lh := handler.NewLoginHandler()
	ad := handler.NewAdminHandler(m)

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
