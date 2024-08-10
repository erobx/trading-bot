package handler

import (
	"fmt"
	"net/http"

	"github.com/erobx/trading-bot/pkg/app/common"
	"github.com/erobx/trading-bot/pkg/view"
)

type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		lh.Post(w, r)
		return
	}
	lh.Get(w, r)
}

func (lh *LoginHandler) Post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Errorf(err.Error())
	}

	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")

	if !common.ValidateLogin(email, password) {
		fmt.Println("invalid login")
		return
	}

	// db check
	_ = email
	_ = password
}

func (lh *LoginHandler) Get(w http.ResponseWriter, r *http.Request) {
	lh.View(w, r)
}

func (lh *LoginHandler) View(w http.ResponseWriter, r *http.Request) {
	view.Login().Render(r.Context(), w)
}
