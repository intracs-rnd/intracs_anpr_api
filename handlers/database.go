package handlers

import (
	"database/sql"
)

func DBConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_intracs_anpr")

	if err != nil {
		return db, err
	}

	return db, nil
}
