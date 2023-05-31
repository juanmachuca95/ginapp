package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConexionSql() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/questions.db")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	err = seedInit(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func seedInit(db *sql.DB) error {
	stmt, err := db.Prepare(QUESTION_TABLE)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
