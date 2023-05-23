package repository

import (
	"database/sql"
	"errors"
	"forum/models"
	"strings"
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
		_, err = tx.Exec("INSERT INTO post_category(post_id, category_name) VALUES (?, ?)", postID, categoryName)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}

func (r *PostSQL) GetPost(id int) (models.Post, error) {
	var p models.Post
	query := `
	SELECT post.id, post.user_id, post.author, post.title, post.content, post.created, post.updated,
		COUNT(CASE WHEN post_rating.islike = 1 THEN 1 END) AS likes, 
		COUNT(CASE WHEN post_rating.islike = -1 THEN 1 END) AS dislikes
	FROM 
		post 
	LEFT JOIN 
		post_rating ON post.id = post_rating.post_id 
	WHERE
		post.id = ?
	GROUP BY post.id;
	`
	row := r.db.QueryRow(query, id)
	err := row.Scan(&p.ID, &p.User_ID, &p.Author, &p.Title, &p.Content, &p.Created, &p.Updated, &p.Likes, &p.Dislikes)
	return p, err
}

// Get all posts with their categories and number of likes, dislikes and comments.
func (r *PostSQL) GetAllPosts() ([]models.Post, error) {
	query := `SELECT
	    post.id,
	    post.user_id,
	    post.author,
	    post.title,
	    post.content,
	    post.created,
	    post.updated,
	    COUNT(DISTINCT comment.id) AS comment_count,
	    COUNT(DISTINCT CASE WHEN pr.islike = 1 THEN pr.user_id || '-' || pr.post_id END) AS like_count,
	    COUNT(DISTINCT CASE WHEN pr.islike = -1 THEN pr.user_id || '-' || pr.post_id END) AS dislike_count
	FROM post
	LEFT JOIN comment ON post.id = comment.post_id
	LEFT JOIN post_rating as pr ON post.id = pr.post_id
	GROUP BY post.id;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	posts := []models.Post{}
	for rows.Next() {
		var post models.Post
		var categories sql.NullString // Use sql.NullString instead of string
		err = rows.Scan(
			&post.ID,
			&post.User_ID,
			&post.Author,
			&post.Title,
			&post.Content,
			&post.Created,
			&post.Updated,
			&post.Comments,
			&post.Likes,
			&post.Dislikes,
			&categories, // Scan as sql.NullString
		)
		if err != nil {
			// Handle the error
		}
		if categories.Valid {
			post.Category = strings.Split(categories.String, ",")
		} else {
			post.Category = []string{} // Set an empty slice for NULL values
		}
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

func (r *PostSQL) DeletePost(user_id, post_id int) error {
	stmt, err := r.db.Prepare("DELETE FROM post WHERE user_id = ? AND id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user_id, post_id)
	return err
}

func (r *PostSQL) LikeDis(rate models.RatePost) error {
	var oldIslike int8

	err := r.db.QueryRow("SELECT islike FROM post_rating WHERE user_id = ? AND post_id = ?", rate.User_ID, rate.Post_ID).Scan(&oldIslike)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return err
		}
	}
	if oldIslike == rate.IsLike {
		stmt, err := r.db.Prepare("DELETE FROM post_rating WHERE user_id = ? AND post_id = ?")
		if err != nil {
			return err
		}
		_, err = stmt.Exec(rate.User_ID, rate.Post_ID)
		return err
	}
	stmt, err := r.db.Prepare(`INSERT INTO post_rating (user_id, post_id, islike) 
	VALUES (?, ?, ?) 
	ON CONFLICT(user_id, post_id) DO UPDATE 
	SET islike = excluded.islike`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(rate.User_ID, rate.Post_ID, rate.IsLike)
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
