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
	row := r.db.QueryRow("SELECT id, user_id, token, expiration_date FROM session WHERE token = ?", token)
	err := row.Scan(&session.ID, &session.UserId, &session.Token, &session.Expiration_date)
	return session, err
}

// func (r *SessionSQL) UpdateSession(s models.Session) error {
// 	stmt, err := r.db.Prepare("UPDATE session SET token = ?, expiration_date = ? WHERE user_id = ?")
// 	if err != nil {
// 		return err
// 	}
// 	_, err = stmt.Exec(s.Token, s.Expiration_date)
// 	return err
// }

func (r *SessionSQL) DeleteSession(user_id int) error {
	stmt, err := r.db.Prepare("DELETE FROM session WHERE user_id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user_id)
	return err
}
