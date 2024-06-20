package handler

import (
	"encoding/json"
	"net/http"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/app/service"
	"github.com/erobx/trading-bot/pkg/view"
)

type DefaultHandler struct {
	svc service.Service
}

func NewDefaultHandler(svc service.Service) *DefaultHandler {
	return &DefaultHandler{
		svc: svc,
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
	g, _ := h.svc.GetGroups(r.Context())
	props := ViewProps{
		Skins: g,
	}
	h.View(w, r, props)
}

func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Success"))
}

type ViewProps struct {
	Skins []model.DisplayGroup
}

func (h *DefaultHandler) View(w http.ResponseWriter, r *http.Request, props ViewProps) {
	view.Page(props.Skins).Render(r.Context(), w)
}

func WriteSkin(skin model.Skin, w http.ResponseWriter) {
	jsonData, err := json.Marshal(skin)
	if err != nil {
		return
	}
	w.Write(jsonData)
}
