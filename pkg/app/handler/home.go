package handler

import (
	"net/http"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/view"
)

type Settings struct {
	isLoggedIn bool
	inv        []model.Skin
}

type DefaultHandler struct {
	market *db.Market
}

func NewDefaultHandler(m *db.Market) *DefaultHandler {
	return &DefaultHandler{
		market: m,
	}
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r)
}

func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	inv, err := h.market.GetInventory("2")
	if err != nil {
		panic(err)
	}

	// call to db for auth or middleware
	h.View(w, r, Settings{
		isLoggedIn: false,
		inv:        inv,
	})
}

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request, c Settings) {
	view.Index(c.isLoggedIn, c.inv).Render(r.Context(), w)
}
