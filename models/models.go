package models

import (
	"time"
)

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
	ID   int8
	Name string
}

type Post struct {
	ID         int
	User_ID    int
	Author     string
	Category   []string
	Title      string
	Content    string
	Created    time.Time
	Updated    time.Time
	Likes      int
	Dislikes   int
	CreatedStr string
	UpdatedStr string
}

type PostCategory struct {
	Post_id     int
	Category_id int8
}

type Comment struct {
	ID       int
	User_ID  int
	Post_ID  int
	Content  string
	Likes    int
	Dislikes int
}

type LikeDis struct {
	ID         int
	User_ID    int
	Post_ID    int
	Comment_ID int
	IsLike     int8
}

type DataList struct {
	User       User
	Posts      []Post
	Categories []string
	Comments   []Comment
}

func (p *Post) TimeToStr() {
	p.CreatedStr = p.Created.Format("02-01-2006 15:04")
	p.UpdatedStr = p.Updated.Format("02-01-2006 15:04")
}
