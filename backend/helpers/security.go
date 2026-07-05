package helpers

import (
	"strings"
	"net/mail"
	"unicode"
	"golang.org/x/crypto/bcrypt"

)

func IsValidEmail(email string) bool {
	email = strings.TrimSpace(email)

	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsValidUserName(username string) bool {
	if len(username) < 3 || len(username) > 20 {
		return false
	}

	runes := []rune(username)

	if !unicode.IsLetter(runes[0]) {
		return false
	}

	for _, r := range runes {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '-') {
			return false
		}
	}

	return true
}

func IsValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 72 {
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSymbol bool

	for _, r := range password {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasNumber = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSymbol = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSymbol
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
