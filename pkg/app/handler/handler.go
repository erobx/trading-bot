package handler

import (
	"encoding/json"
	"net/http"

	"github.com/erobx/trading-bot/pkg/app/model"
	"github.com/erobx/trading-bot/pkg/app/service"
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
	skin, err := h.svc.GetSkin(r.Context(), "Redline", "Factory New")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	WriteSkin(skin, w)
}

func (h *DefaultHandler) Post(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Success"))
}

func WriteSkin(skin model.Skin, w http.ResponseWriter) {
	jsonData, err := json.Marshal(skin)
	if err != nil {
		return
	}
	w.Write(jsonData)
}
