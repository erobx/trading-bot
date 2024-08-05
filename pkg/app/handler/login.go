package handler

import "net/http"

type LoginHandler struct{}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

func (lh *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

func (lh *LoginHandler) View(w http.ResponseWriter, r *http.Request) {
}
