package app

import (
	"belajar-go-rest-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/go_rest_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
