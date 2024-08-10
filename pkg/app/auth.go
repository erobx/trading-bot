package app

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

type Auth struct {
	next     http.Handler
	username string
	password string
}

func NewAuth(n http.Handler, username, password string) http.Handler {
	auth := Auth{
		next:     n,
		username: username,
		password: password,
	}
	return auth
}

func (a Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if ok {
		usernameHash := sha256.Sum256([]byte(username))
		passwordHash := sha256.Sum256([]byte(password))
		expectedUsername := sha256.Sum256([]byte(a.username))
		expectedPasword := sha256.Sum256([]byte(a.password))

		usernameMatch := (subtle.ConstantTimeCompare(usernameHash[:], expectedUsername[:]) == 1)
		passwordMatch := (subtle.ConstantTimeCompare(passwordHash[:], expectedPasword[:]) == 1)

		if usernameMatch && passwordMatch {
			a.next.ServeHTTP(w, r)
			return
		}
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
