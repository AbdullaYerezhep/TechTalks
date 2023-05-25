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
	ErrUserExists      = errors.New("username or password already exists")
	// post action errors
	ErrPostContent = errors.New("content size must be more than 10 and less ")
	ErrPostTitle   = errors.New("title size must be more than 5 and less than 35")
	// comment action errors
	ErrPermission     = errors.New("permission denied")
	ErrCommentContent = errors.New("content size must be more than 1 and less 100")
)
