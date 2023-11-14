package database

import (
	"database/sql"
	"fmt"
	"intracs_anpr_api/internal/env"
	"time"

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

	src := user + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=true&loc=Local"
	sqlDB, err := sql.Open(driver, src)
	if err != nil {
		fmt.Println("SQL connection have problem", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		fmt.Println("Cannot connect to database", err)
	}

	return db
}
