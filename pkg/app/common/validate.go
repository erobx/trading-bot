package common

import (
	"net/mail"

	"github.com/erobx/trading-bot/pkg/db"
)

func ValidateNewUser(m *db.Market, username, email, password string) bool {
	return validateUsername(username) && validateEmail(m, email) && validatePass(password)
}

func validateUsername(name string) bool {
	if len(name) <= 0 {
		return false
	}
	return true
}

func validateEmail(m *db.Market, email string) bool {
	//m.CheckEmail()
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validatePass(password string) bool {
	if len(password) <= 0 {
		return false
	}
	return true
}
