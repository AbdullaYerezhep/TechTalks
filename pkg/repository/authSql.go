package repository

import (
	"database/sql"
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

func (r *AuthSQL) CreateUser(u models.User) error {
	statement, err := r.db.Prepare("INSERT INTO users (username, email, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(u.Name, u.Email, u.Password)
	return err
}

func (r *AuthSQL) GetUser(name string) (models.User, error) {
	var u models.User
	err := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", name).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	return u, err
}

func (r *AuthSQL) GetUserByID(id int) (models.User, error) {
	var u models.User
	err := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE id = ?", id).Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	return u, err
}

func (r *AuthSQL) GetUserByToken(token string) (models.User, error) {
	var s models.Session
	err := r.db.QueryRow("SELECT id, user_id, token, expiration_date FROM session WHERE token = ?", token).Scan(&s.ID, &s.UserId, &s.Token, &s.Expiration_date)
	if err != nil {
		return models.User{}, err
	}
	return r.GetUserByID(s.UserId)
}
