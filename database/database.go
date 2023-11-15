package database

import (
	"database/sql"
	"fmt"
	"intracs_anpr_api/internal/env"
	"log"
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

	if port == "" {
		port = "3306"
	}

	if password != "" {
		password = ":" + password
	}

	src := user + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=true&loc=Local"
	sqlDB, err := sql.Open(driver, src)
	if err != nil {
		log.Fatal("SQL connection have problem", err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{Conn: sqlDB}), &gorm.Config{
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Jakarta")
			return time.Now().In(ti)
		},
	})
	if err != nil {
		log.Fatal("Can't connect to database, cause:", err)
	}

	// Remove ONLY_FULL_GROUP_BY on sql_mode
	db.Raw("SET PERSIST sql_mode=(SELECT REPLACE(@@sql_mode,'ONLY_FULL_GROUP_BY',''))")
	if db.Error != nil {
		log.Fatal("Can't remove ONLY_FULL_GROUP_BY on sql mode, cause:", err)
	}

	fmt.Printf("Connected to database %s:%s\n", host, port)
	return db
}
