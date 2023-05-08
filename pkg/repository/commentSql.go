package repository

import (
	"database/sql"
	"forum/models"
)

type CommentSQL struct {
	db *sql.DB
}

func NewCommentSQL(db *sql.DB) *CommentSQL {
	return &CommentSQL{db: db}
}

func (r *CommentSQL) AddComment(com models.Comment) error {
	stmt, err := r.db.Prepare("INSERT INTO comment (user_id, post_id, content) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(com.User_ID, com.Post_ID, com.Content)
	return err
}

func (r *CommentSQL) GetComment(id int) (models.Comment, error) {
	var com models.Comment
	row := r.db.QueryRow("SELECT id, user_id, post_id, content FROM comment WHERE id = ?", id)
	err := row.Scan(&com.ID, &com.User_ID, &com.Post_ID, &com.Content)
	return com, err
}

func (r *CommentSQL) GetPostComments(post_id int) ([]models.Comment, error) {
	var comments []models.Comment
	row, err := r.db.Query("SELECT id, user_id, post_id, content FROM comment WHERE post_id = ?", post_id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var com models.Comment
		err := row.Scan(&com.ID, &com.User_ID, &com.Post_ID, &com.Content)
		if err != nil {
			return nil, err
		}
		comments = append(comments, com)
	}
	return comments, nil
}

func (r *CommentSQL) UpdateComment(com models.Comment) error {
	stmt, err := r.db.Prepare("UPDATE comment SET content = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(com.Content, com.ID)
	return err
}

func (r *CommentSQL) DeleteComment(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM comment WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
