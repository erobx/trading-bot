package handler

import (
	"net/http"

	"github.com/erobx/trading-bot/pkg/db"
	"github.com/erobx/trading-bot/pkg/view"
)

type ModalHandler struct {
	market *db.Market
}

func NewModalHandler(m *db.Market) *ModalHandler {
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

	gid := r.FormValue("gid")
	sid := r.FormValue("sid")

	err := h.market.AddSkinToGroup(gid, sid)
	if err != nil {
		panic(err)
	}

	g, _ := h.market.GetActiveGroups()
	h.View(w, r, ViewProps{
		Groups: g,
	})
}

func (h *ModalHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	view.Groups(props.Groups).Render(r.Context(), w)
}
