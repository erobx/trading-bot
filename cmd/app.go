package main

import (
	"encoding/json"
	"net/http"
)

type App struct {
	user *User
}

func NewApp(user *User) *App {
	return &App{
		user: user,
	}
}

func (s *App) start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.handleIndex)
	mux.HandleFunc("/addskin", s.addSkin)

	http.ListenAndServe(":3000", mux)
}

func (s *App) handleIndex(w http.ResponseWriter, r *http.Request) {
	skin, err := s.user.GetSkin("Redline", "Factory New", 32)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	WriteSkin(skin, w)
}

func (s *App) addSkin(w http.ResponseWriter, r *http.Request) {
	if err := s.user.AddSkin("Redline", "Factory New", 32); err != nil {
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func WriteSkin(skin Skin, w http.ResponseWriter) {
	jsonData, err := json.Marshal(skin)
	if err != nil {
		return
	}
	w.Write(jsonData)
}
