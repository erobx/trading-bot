package handler

import (
	"net/http"
	"time"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/view"
)

type TradeupsHandler struct {
	market *db.Market
	Now    func() time.Time
}

func NewTradeupsHandler(m *db.Market, now func() time.Time) *TradeupsHandler {
	return &TradeupsHandler{
		market: m,
		Now:    now,
	}
}

func (h *TradeupsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r)
}

func (h *TradeupsHandler) Get(w http.ResponseWriter, r *http.Request) {
	t, _ := h.market.GetActiveTradeups()
	h.View(w, r, ViewProps{
		Tradeups: t,
	})
}

type ViewProps struct {
	Tradeups []model.DisplayTrade
}

func (h *TradeupsHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	view.Tradeups(props.Tradeups, h.Now(), false).Render(r.Context(), w)
}
