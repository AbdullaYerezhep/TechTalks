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
	UserId          int
	Token           string
	Expiration_date time.Time
}

type Category struct {
	ID   int8
	Name string
}

type Post struct {
	ID        int `json:"id"`
	User_ID   int
	Author    string   `json:"author"`
	Category  []string `json:"categories"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	CreatedAt time.Time
	UpdatedAt *time.Time
	Comments  int
	Likes     int
	Dislikes  int
	Created   string
	Updated   *string
}

type PostCategory struct {
	Post_id     int
	Category_id int8
}

type Comment struct {
	ID       int `json:"id"`
	User_ID  int
	Author   string `json:"username"`
	Post_ID  int    `json:"post_id"`
	Content  string `json:"content"`
	Created  string
	Updated  *string
	Likes    int
	Dislikes int
}

type RatePost struct {
	User_ID int
	Post_ID int  `json:"post_id"`
	IsLike  int8 `json:"islike"`
}

type RateComment struct {
	User_ID    int
	Post_ID    int  `json:"post_id"`
	Comment_ID int  `json:"comment_id"`
	IsLike     int8 `json:"islike"`
}

type HomePage struct {
	User
	Posts      []Post
	Categories []string
}

type AddPostPage struct {
	User
	Categories []string
}

type PostPageData struct {
	*User
	Post
	Comments   []Comment
	Categories []string
}

// func (p *Post) TimeToStr() {
// 	p.CreatedStr = p.Created.Format("02-01-2006 15:04")
// 	if p.Updated != nil {
// 		*p.UpdatedStr = p.Updated.Format("02-01-2006 15:04")
// 	}
// }
