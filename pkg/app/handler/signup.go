package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/argon2id"
	"github.com/erobx/trading-bot/pkg/app/common"
	"github.com/erobx/trading-bot/pkg/app/model"
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

func (sh *SignupHandler) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	if !common.ValidateNewUser(sh.market, username, email, password) {
		log.Println(fmt.Errorf("invalid user"))
		return
	}

	user := model.User{
		Username: username,
		Email:    email,
	}
	fmt.Println("Created new user", user.Username, user.Email)

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
}

func (sh *SignupHandler) View(w http.ResponseWriter, r *http.Request) {
	view.Signup().Render(r.Context(), w)
}
