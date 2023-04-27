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

type Category struct {
	ID   int
	Name string
}

type Post struct {
	ID       int
	User_ID  int
	Category string
	Title    string
	Content  string
	Created  time.Time
	Updated  time.Time
}

type PostCategory struct {
	Post_id     int
	Category_id int
}

type Comment struct {
	ID      int
	User_ID int
	Post_ID int
	Content string
}

type LikeDis struct {
	ID         int
	User_ID    int
	Post_ID    int
	Comment_ID int
	IsLike     int8
}
