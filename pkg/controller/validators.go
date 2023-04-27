package controller

import (
	"errors"
	"forum/models"
	"net/url"
	"regexp"
	"unicode"
)

const (
	nameRegExp  = `^[a-zA-Z0-9_.'-]{3,15}$`
	emailRegExp = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{3,29}$`
)

func decodeForm(form url.Values) (models.User, error) {
	var u models.User
	name, okName := form["username"]
	email, okEmail := form["email"]
	password, okPass := form["password"]
	confirmPass, okConfirm := form["confirmPass"]
	if !okName || !okEmail || !okPass || !okConfirm {
		return u, errors.New("form modified")
	}
	if err := validateNewUserData(name[0], email[0], password[0], confirmPass[0]); err != nil {
		return u, err
	}
	u.Name = name[0]
	u.Email = email[0]
	u.Password = password[0]

	return u, nil
}

func validateNewUserData(name, email, password, confirmPass string) error {
	if !isAscii(name) || !isAscii(password) || !isAscii(password) {
		return errors.New("non-ascii character")
	}
	if !isValidName(name) {
		return errors.New("invalid username")
	}
	if !isValidEmail(email) {
		return errors.New("invalid email")
	}
	if !isValidPassword(password, confirmPass) {
		return errors.New("invalid password")
	}
	return nil
}

func isAscii(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] > unicode.MaxASCII {
			return false
		}
	}
	return true
}

func isValidName(name string) bool {
	re := regexp.MustCompile(nameRegExp)
	return re.MatchString(name)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(emailRegExp)
	return re.MatchString(email)
}

func isValidPassword(password, confirmPass string) bool {
	// re := regexp.MustCompile(passRegExp)
	// return re.MatchString(password)
	if len(password) < 8 || len(password) > 25 {
		return false
	}
	return password == confirmPass
}
