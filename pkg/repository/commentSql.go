package repository

import (
	"database/sql"
	"errors"
	"forum/models"
)

type CommentSQL struct {
	db *sql.DB
}

func NewCommentSQL(db *sql.DB) *CommentSQL {
	return &CommentSQL{db: db}
}

func (r *CommentSQL) AddComment(com models.Comment) error {
	err := r.db.QueryRow("SELECT username FROM users WHERE id = ?", com.User_ID).Scan(&com.Author)
	if err != nil {
		return err
	}
	stmt, err := r.db.Prepare("INSERT INTO comment (user_id, username, post_id, content, created) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(com.User_ID, com.Author, com.Post_ID, com.Content, com.Created)
	return err
}

func (r *CommentSQL) GetComment(id int) (models.Comment, error) {
	var com models.Comment
	row := r.db.QueryRow("SELECT id, user_id, post_id, content FROM comment WHERE id = ?", id)
	err := row.Scan(&com.ID, &com.User_ID, &com.Post_ID, &com.Content)
	return com, err
}

func (r *CommentSQL) GetPostComments(post_id int) ([]models.Comment, error) {
	query := `SELECT 
		comment.id, 
		comment.user_id, 
		comment.username, 
		comment.post_id, 
		comment.content, 
		comment.created, 
		comment.updated, 
	    COUNT(comment_rating.islike) as likes,
	    COUNT(CASE WHEN comment_rating.islike = -1 THEN 1 END) as dislikes
	FROM 
		comment
	LEFT JOIN comment_rating ON comment.id = comment_rating.comment_id
	WHERE comment.post_id = ?
	GROUP BY comment.id, comment.user_id, comment.username, comment.post_id, comment.content, comment.created, comment.updated
	ORDER BY comment.created ASC
	`
	var comments []models.Comment
	row, err := r.db.Query(query, post_id)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var com models.Comment
		err := row.Scan(&com.ID, &com.User_ID, &com.Author, &com.Post_ID, &com.Content, &com.Created, &com.Updated, &com.Likes, &com.Dislikes)
		if err != nil {
			return nil, err
		}
		comments = append(comments, com)
	}
	return comments, nil
}

func (r *CommentSQL) UpdateComment(com models.Comment) error {
	stmt, err := r.db.Prepare("UPDATE comment SET content = ?, updated = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(com.Content, com.Updated, com.ID)
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

func (r *CommentSQL) RateComment(rate models.RateComment) error {
	var oldIslike int8

	err := r.db.QueryRow("SELECT islike FROM comment_rating WHERE comment_id = ? AND user_id = ?", rate.Comment_ID, rate.User_ID).Scan(&oldIslike)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}

	if oldIslike == rate.IsLike {
		stmt, err := r.db.Prepare("DELETE FROM comment_rating WHERE comment_id = ? AND user_id = ?")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(rate.Comment_ID, rate.User_ID)
		return err
	}

	stmt, err := r.db.Prepare(`
	INSERT INTO 
		comment_rating (comment_id, user_id, islike) 
	VALUES (?, ?, ?) 
	ON CONFLICT(comment_id, user_id) 
	DO UPDATE 
		SET islike = excluded.islike`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(rate.Comment_ID, rate.User_ID, rate.IsLike)
	return err
}
