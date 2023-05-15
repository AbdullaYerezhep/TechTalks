package repository

import (
	"database/sql"
)

type CategorySQL struct {
	db *sql.DB
}

func NewCategorySQL(db *sql.DB) *CategorySQL {
	return &CategorySQL{db: db}
}

func (r *CategorySQL) GetCategories() ([]string, error) {
	rows, err := r.db.Query("SELECT name FROM category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	categories := []string{}
	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}
