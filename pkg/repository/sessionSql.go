package repository

import (
	"database/sql"
	"forum/models"
)

type SessionSQL struct {
	db *sql.DB
}

func NewSessionSQL(db *sql.DB) *SessionSQL {
	return &SessionSQL{
		db: db,
	}
}

func (r *SessionSQL) CreateSession(s models.Session) error {
	statement, err := r.db.Prepare("INSERT INTO session (user_id, token, expiration_date) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = statement.Exec(s.UserId, s.Token, s.Expiration_date)
	return err

}

func (r *SessionSQL) GetSession(token string) (models.Session, error) {
	var session models.Session
	err := r.db.QueryRow("SELECT id, user_id, token, expiration_date FROM session WHERE token = ?", token).Scan(&session.ID, &session.UserId, &session.Token, &session.Expiration_date)
	return session, err
}

func (r *SessionSQL) DeleteSession(id int) error {
	r.db.QueryRow("DELETE id, user_id token, expiration_date FROM session WHERE user_id = ?", id)
	return nil
}
