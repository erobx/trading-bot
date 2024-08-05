package handler

import (
	"net/http"
	"time"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/view"
)

type GroupsHandler struct {
	market *db.Market
	Now    func() time.Time
}

func NewGroupsHandler(m *db.Market, now func() time.Time) *GroupsHandler {
	return &GroupsHandler{
		market: m,
		Now:    now,
	}
}

func (h *GroupsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Get(w, r)
}

func (h *GroupsHandler) Get(w http.ResponseWriter, r *http.Request) {
	g, _ := h.market.GetActiveGroups()
	h.View(w, r, ViewProps{
		Groups: g,
	})
}

type ViewProps struct {
	Groups []model.DisplayGroup
}

func (h *GroupsHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	view.Groups(props.Groups, h.Now(), false).Render(r.Context(), w)
}
