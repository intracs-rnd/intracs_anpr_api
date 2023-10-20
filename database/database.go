package database

import (
	"database/sql"
	"fmt"
	"intracs_anpr_api/internal/env"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabase() *gorm.DB {
	driver := env.Get("DB_DRIVER")
	name := env.Get("DB_NAME")
	host := env.Get("DB_HOST")
	port := env.Get("DB_PORT")
	user := env.Get("DB_USER")
	password := env.Get("DB_PASSWORD")

	src := user + password + "@tcp(" + host + ":" + port + ")/" + name
	sqlDB, err := sql.Open(driver, src)
	if err != nil {
		fmt.Println("SQL connection have problem", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{})
	if err != nil {
		fmt.Println("Cannot connect to database", err)
	}

	return db
}
