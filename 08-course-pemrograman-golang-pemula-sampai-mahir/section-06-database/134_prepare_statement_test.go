package belajar_golang_database

import (
	"database/sql"
	"testing"
	"time"
)

func PrepareStatementGetConnection(driver, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(20 * time.Minute)

	return db
}

func TestPrepareStatement(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := PrepareStatementGetConnection(driver, source)
	defer db.Close()
}
