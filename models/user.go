package models

import "time"

type User struct {
	ID       int
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	ID              int
	UserId          int
	Token           string
	Expiration_date time.Time
}

type Post struct {
	ID       int
	UID      int // author id
	Category string
	Title    string
	Content  string
	Created  time.Time
	Updated  time.Time
}
