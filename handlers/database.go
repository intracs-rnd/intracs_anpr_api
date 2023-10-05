package handlers

import (
	"database/sql"
	. "intracs_anpr_api/configs"
)

var driver string
var name string
var host string
var port string
var user string
var password string = ""
var dataSource string

func init() {
	driver = GetEnv("DB_DRIVER")
	name = GetEnv("DB_NAME")
	host = GetEnv("DB_HOST")
	port = GetEnv("DB_PORT")
	user = GetEnv("DB_USER")

	if GetEnv("DB_PASSWORD") != "" {
		password = ":" + GetEnv("DB_PASSWORD")
	}

	dataSource = user + password + "@tcp(" + host + ":" + port + ")/" + name
	println(dataSource)
}

func DBConnect() (*sql.DB, error) {
	// db, err := sql.Open("mysql", "root@tcp(localhost:3306)/db_intracs_anpr")
	db, err := sql.Open(driver, dataSource)

	if err != nil {
		return db, err
	}
	// defer db.Close()

	return db, nil
}
