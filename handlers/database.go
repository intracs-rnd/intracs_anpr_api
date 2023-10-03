package handlers

import (
	"database/sql"
)

func DBConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_intracs_anpr")
	db.SetMaxOpenConns(0)

	if err != nil {
		return db, err
	}
	// defer db.Close()

	return db, nil
}
