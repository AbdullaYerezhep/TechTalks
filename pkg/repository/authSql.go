package repository

import (
	"database/sql"
	"fmt"
	"forum/models"
)

type AuthSQL struct {
	db *sql.DB
}

func NewAuthSQL(db *sql.DB) *AuthSQL {
	return &AuthSQL{
		db: db,
	}
}

func (a *AuthSQL) CreateUser(u *models.User) error {
	statement, err := a.db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}
	fmt.Println("User has been created!")
	return nil
}

func (a *AuthSQL) GetByName(name string) (*models.User, error) {
	var u *models.User
	err := a.db.QueryRow("SELECT id, email, password FROM users WHERE username = ?", name).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return nil, err
	}
	return u, nil
}
