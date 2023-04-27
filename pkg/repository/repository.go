package repository

import (
	"database/sql"
	"forum/models"
)

type Authorization interface {
	CreateUser(user models.User) error
	GetUser(name string) (models.User, error)
	GetUserByToken(token string) (models.User, error)
}

type Session interface {
	CreateSession(models.Session) error
	GetSession(token string) (models.Session, error)
	DeleteSession(id int) error
}

type Post interface {
	Create(user_id int, title, content string) error
	Get(title string) (models.Post, error)
	Delete(id int) error
	Update(id int, newTitle, newContent string) error
}

type Repository struct {
	Authorization
	Session
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
		Session:       NewSessionSQL(db),
	}
}
