package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

type DefaultHandler struct {
	svc Service
}

func NewDefaultHandler(svc Service) *DefaultHandler {
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
	// temp, _ := decimal.NewFromString("123.123")
	// price := dbDecimal(temp)
	// err := h.svc.AddSkin(r.Context(), "Redline", "Factory New", price)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	h.addDummyData()
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Success"))
}

func WriteSkin(skin Skin, w http.ResponseWriter) {
	jsonData, err := json.Marshal(skin)
	if err != nil {
		return
	}
	w.Write(jsonData)
}

func (h *DefaultHandler) addDummyData() {
	names := []string{"Redline", "Water Elemental", "Block9"}
	wears := [5]string{"Factory New", "Minimal Wear", "Field-Tested", "Well-Worn", "Battle-Scarred"}
	prices := randomPrices()

	for i := 0; i < 40; i++ {
		nameIndex := rand.Intn(len(names))
		wearIndex := rand.Intn(len(wears))
		priceIndex := rand.Intn(len(prices))

		skin := Skin{
			Name:  names[nameIndex],
			Wear:  wears[wearIndex],
			Price: prices[priceIndex],
		}
		err := h.svc.AddSkin(context.TODO(), skin.Name, skin.Wear, skin.Price)
		if err != nil {
			panic(err)
		}
	}
}
