package repository

import (
	"database/sql"
	"io/ioutil"

	_ "github.com/mattn/go-sqlite3"
)

func NewDB(dbtype, dbname string) (*sql.DB, error) {
	db, err := sql.Open(dbtype, dbname)
	if err != nil {
		return nil, err
	}
	if err = migrate(db); err != nil {
		return nil, err
	}
	return db, err
}

func migrate(db *sql.DB) error {
	fileByte, err := ioutil.ReadFile("migrations.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(fileByte))
	if err != nil {
		return err
	}
	return nil
}
