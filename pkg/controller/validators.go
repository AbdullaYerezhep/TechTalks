package controller

import (
	"errors"
	"regexp"
	"unicode"
)

const (
	nameRegExp  = `^[a-zA-Z0-9_.'-]{3,15}$`
	emailRegExp = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{3,29}$`
)

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
	return isValidPassword(password, confirmPass)
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

func isValidPassword(password, confirmPass string) error {
	var err error
	// match1, _ := regexp.MatchString("[A-Z]", password)
	// match2, _ := regexp.MatchString("[a-z]", password)
	// match3, _ := regexp.MatchString("[~!@#$%^&*_()-+={[}]|\\:;\"'<,>.?/]", password)

	// if !match1 || !match2 || !match3 {
	// 	return errors.New("invalid password")
	// }

	if confirmPass != password {
		err = errors.New("confirmation failed")
	}
	return err
}
