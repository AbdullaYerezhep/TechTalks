package repository

import (
	"database/sql"
	"forum/models"
)

type Authorization interface {
	CreateUser(user *models.User) error
	GetId()
	GetByName(name string) (*models.User, error)
}

type Repository struct {
	Authorization
}

func New(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQL(db),
	}
}
