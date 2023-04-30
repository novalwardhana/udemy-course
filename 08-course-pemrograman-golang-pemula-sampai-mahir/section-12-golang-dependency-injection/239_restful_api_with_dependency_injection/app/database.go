package app

import (
	"database/sql"
	"restful-api/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func Database() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api")
	helper.CategoryPanicIfError(err)

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(2)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db
}
