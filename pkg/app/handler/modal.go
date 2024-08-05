package handler

import (
	"net/http"
	"strconv"

	"github.com/erobx/trading-bot/pkg/app/model"
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

	groupId := r.FormValue("gid")
	skinId := r.FormValue("sid")

	err := h.market.AddSkinToGroup(groupId, skinId)
	if err != nil {
		panic(err)
	}

	// get group that changed
	g, err := h.market.GetChangedGroup(groupId)
	if err != nil {
		panic(err)
	}

	h.View(w, r, GroupProps{
		ID:    strconv.Itoa(g.GroupId),
		Tier:  g.Tier,
		Skins: g.Skins,
	})
}

type GroupProps struct {
	ID    string
	Tier  string
	Skins []model.Skin
}

func (h *ModalHandler) View(w http.ResponseWriter, r *http.Request, props GroupProps) {
	view.Group(props.ID, props.Tier, props.Skins).Render(r.Context(), w)
}
