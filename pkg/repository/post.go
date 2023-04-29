package repository

import (
	"database/sql"
	"forum/models"
)

type PostSQL struct {
	db *sql.DB
}

func NewPostSQL(db *sql.DB) *PostSQL {
	return &PostSQL{
		db: db,
	}
}

func (r *PostSQL) CreatePost(p models.Post) error {
	stmt, err := r.db.Prepare("INSERT INTO post (user_id, author, title, content, created, updated) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(p.User_ID, p.Author, p.Title, p.Content, p.Created, p.Updated)
	return err
}

func (r *PostSQL) GetPost(id int) (models.Post, error) {
	var p models.Post
	row := r.db.QueryRow("SELECT id, user_id, author, title, content, created, updated FROM post WHERE id = ?", id)
	err := row.Scan(&p.ID, &p.User_ID, &p.Author, &p.Title, &p.Content, &p.Created, &p.Updated)
	return p, err
}

func (r *PostSQL) GetAllPosts() ([]models.Post, error) {
	rows, err := r.db.Query("SELECT * FROM post ORDER BY created DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.User_ID, &post.Author, &post.Title, &post.Content, &post.Created, &post.Updated)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, err
}

func (r *PostSQL) UpdatePost(p models.Post) error {
	stmt, err := r.db.Prepare("UPDATE post SET title = ?, content = ?, updated = ? where id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(p.Title, p.Content, p.Updated, p.ID)
	return err
}

func (r *PostSQL) DeletePost(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM post WHERE id = ?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id)
	return err
}
