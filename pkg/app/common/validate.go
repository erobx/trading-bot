package common

import (
	"net/mail"

	"github.com/erobx/trading-bot/pkg/db"
)

func ValidateNewUser(m *db.Market, username, email, password string) bool {
	return validateUsername(username) && validateEmail(email) && validatePass(password)
}

func ValidateLogin(email, password string) bool {
	return validateEmail(email) && validatePass(password)
}

func validateUsername(name string) bool {
	if len(name) <= 0 {
		return false
	}
	return true
}

func validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func validatePass(password string) bool {
	if len(password) <= 0 {
		return false
	}
	return true
}
