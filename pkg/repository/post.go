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
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO post (user_id, author, title, content, created, updated) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.User_ID, p.Author, p.Title, p.Content, p.Created, p.Updated)
	if err != nil {
		tx.Rollback()
		return err
	}

	postID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return err
	}
	for _, categoryName := range p.Category {
		var categoryID int8
		err = r.db.QueryRow("SELECT id FROM category WHERE name = ?", categoryName).Scan(&categoryID)
		if err != nil {
			tx.Rollback()
			return err
		}
		_, err = tx.Exec("INSERT INTO post_category(post_id, category_id) VALUES (?, ?)", postID, categoryID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
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
		post.TimeToStr()
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *PostSQL) UpdatePost(p models.Post) error {
	stmt, err := r.db.Prepare("UPDATE post SET title = ?, content = ?, updated = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Title, p.Content, p.Updated, p.ID)
	return err
}

func (r *PostSQL) DeletePost(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM post WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}

// func (r *PostSQL) GetPostsByCategory(category string) ([]models.Post, error) {
// 	rows, err := r.db.Query("SELECT * FROM post WHERE category = ? ORDER BY created DESC", category)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	posts := []models.Post{}
// 	for rows.Next() {
// 		var post models.Post
// 		err := rows.Scan(&post.ID, &post.User_ID, &post.Author, &post.Title, &post.Content, &post.Created, &post.Updated)
// 		if err != nil {
// 			return nil, err
// 		}
// 		post.TimeToStr()
// 		posts = append(posts, post)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}
// 	return posts, err
// }
