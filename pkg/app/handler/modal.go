package handler

import (
	"net/http"
	"strconv"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/view"
)

type ModalHandler struct {
	market *db.Db
}

func NewModalHandler(m *db.Db) *ModalHandler {
	return &ModalHandler{
		market: m,
	}
}

func (h *ModalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.Post(w, r)
		return
	}
}

func (h *ModalHandler) Post(w http.ResponseWriter, r *http.Request) {
	// add skin to approriate group
	r.ParseForm()

	tradeupId := r.FormValue("tid")
	skinId := r.FormValue("sid")

	err := h.market.AddSkinToTradeup(tradeupId, skinId)
	if err != nil {
		panic(err)
	}

	// get group that changed
	t, err := h.market.GetChangedTradeup(tradeupId)
	if err != nil {
		panic(err)
	}

	h.View(w, r, TradeupProps{
		ID:    strconv.Itoa(t.TradeId),
		Tier:  t.Tier,
		Skins: t.Skins,
	})
}

type TradeupProps struct {
	ID    string
	Tier  string
	Skins []model.Skin
}

func (h *ModalHandler) View(w http.ResponseWriter, r *http.Request, props TradeupProps) {
	view.Tradeup(props.ID, props.Tier, props.Skins).Render(r.Context(), w)
}
