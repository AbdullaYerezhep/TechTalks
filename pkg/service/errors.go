package service

import "errors"

var (
	// authorization errors
	ErrInvalidName     = errors.New("invalid username")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
	ErrAscii           = errors.New("non-ascii character")
	ErrUserNotFound    = errors.New("user not found")
	ErrWrongPassword   = errors.New("wrong password")
	// post action errors
	// comment action errors
	ErrPermission = errors.New("permission denied")
)
