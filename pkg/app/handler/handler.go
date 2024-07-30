package handler

import (
	"net/http"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/view"
)

type DefaultHandler struct {
	market *db.Market
}

func NewDefaultHandler(m *db.Market) *DefaultHandler {
	return &DefaultHandler{
		market: m,
	}
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.Post(w, r)
		return
	}
	h.Get(w, r)
}

func (h *DefaultHandler) Get(w http.ResponseWriter, r *http.Request) {
	g, _ := h.market.GetActiveGroups()
	h.View(w, r, ViewProps{
		Groups: g,
	})
}

func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	g, _ := h.market.GetActiveGroups()
	h.View(w, r, ViewProps{
		Groups: g,
	})
}

type ViewProps struct {
	Groups []model.DisplayGroup
}

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	view.Groups(props.Groups).Render(r.Context(), w)
}
