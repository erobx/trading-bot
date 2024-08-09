package handler

import (
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/view"
)

type SignupHandler struct {
	market *db.Market
}

func NewSignupHandler(m *db.Market) *SignupHandler {
	return &SignupHandler{
		market: m,
	}
}

func (sh *SignupHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		sh.Post(w, r)
		return
	}
	sh.Get(w, r)
}

// create new user
func (sh *SignupHandler) Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	//fmt.Println("Username:", username)
	//fmt.Println("Email:", email)
	//fmt.Println("Pass:", password)

	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}

	err = sh.market.AddUser(username, email, hash)
	if err != nil {
		log.Fatal(err)
	}
}

func (sh *SignupHandler) Get(w http.ResponseWriter, r *http.Request) {
	sh.View(w, r)
	return
}

func (sh *SignupHandler) View(w http.ResponseWriter, r *http.Request) {
	view.Signup().Render(r.Context(), w)
}
