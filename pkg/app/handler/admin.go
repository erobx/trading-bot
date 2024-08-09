package handler

import (
	"net/http"

	"github.com/erobx/trading-bot/pkg/db"
)

type AdminHandler struct {
	market *db.Market
}

func NewAdminHandler(m *db.Market) *AdminHandler {
	return &AdminHandler{
		market: m,
	}
}

func (ad *AdminHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		ad.Post(w, r)
		return
	}
	ad.Get(w, r)
}

func (ad *AdminHandler) Post(w http.ResponseWriter, r *http.Request) {
	return
}

func (ad *AdminHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Dashboard"))
	return
}

func (ad *AdminHandler) View(w http.ResponseWriter, r *http.Request) {
	return

}
