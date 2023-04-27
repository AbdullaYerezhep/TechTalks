package repository

import (
	"database/sql"
	"forum/models"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(name string) (models.User, error)
	GetUserByID(id int) (models.User, error)
	GetUserByToken(token string) (models.User, error)
}

type Session interface {
	CreateSession(models.Session) error
	GetSession(token string) (models.Session, error)
	DeleteSession(id int) error
}

type Post interface {
	CreatePost(p models.Post) error
	GetPost(id int) (models.Post, error)
	GetAllPosts() ([]models.Post, error)
	UpdatePost(p models.Post) error
	DeletePost(id int) error
}

type Repository struct {
	Authorization
	Session
	Post
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
		Session:       NewSessionSQL(db),
		Post:          NewPostSQL(db),
	}
}
