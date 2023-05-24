package repository

import (
	"database/sql"
	"forum/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(name string) (models.User, error)
	GetUserByID(id int) (models.User, error)
}

type Session interface {
	CreateSession(models.Session) error
	GetSession(token string) (models.Session, error)
	DeleteSession(user_id int) error
}

type Post interface {
	CreatePost(p models.Post) error
	GetPost(id int) (models.Post, error)
	GetAllPosts() ([]models.Post, error)
	GetTopPostsByLikes() ([]models.Post, error)
	UpdatePost(p models.Post) error
	DeletePost(user_id, post_id int) error
	LikeDis(models.RatePost) error
}

type Category interface {
	GetCategories() ([]string, error)
}

type Comment interface {
	AddComment(models.Comment) error
	GetComment(id int) (models.Comment, error)
	GetPostComments(id int) ([]models.Comment, error)
	UpdateComment(models.Comment) error
	DeleteComment(id int) error
	RateComment(models.RateComment) error
}

type Repository struct {
	Authorization
	Session
	Post
	Category
	Comment
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
		Session:       NewSessionSQL(db),
		Post:          NewPostSQL(db),
		Category:      NewCategorySQL(db),
		Comment:       NewCommentSQL(db),
	}
}
