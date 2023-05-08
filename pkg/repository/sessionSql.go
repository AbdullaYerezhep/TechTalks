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
	stmt, err := r.db.Prepare(`INSERT INTO session (user_id, token, expiration_date) VALUES (?, ?, ?) 
	ON CONFLICT(user_id) DO UPDATE SET token = excluded.token, expiration_date = excluded.expiration_date`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(s.UserId, s.Token, s.Expiration_date)
	return err
}

func (r *SessionSQL) GetSession(token string) (models.Session, error) {
	var session models.Session
	row := r.db.QueryRow("SELECT user_id, token, expiration_date FROM session WHERE token = ?", token)
	err := row.Scan(&session.UserId, &session.Token, &session.Expiration_date)
	return session, err
}

func (r *SessionSQL) DeleteSession(user_id int) error {
	stmt, err := r.db.Prepare("DELETE FROM session WHERE user_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user_id)
	return err
}
