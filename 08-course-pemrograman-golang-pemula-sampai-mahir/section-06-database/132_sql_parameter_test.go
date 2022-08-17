package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func SQLParameterGetConnection(driver, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(20 * time.Second)

	return db
}

func TestSQLParameterSelect(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := SQLParameterGetConnection(driver, source)
	defer db.Close()

	ctx := context.Background()
	username := "noval'; #"
	password := "secret1231"
	sqlQuery := `select username, password from user where username = ? and password = ? limit 1`

	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username, password string
		rows.Scan(&username, &password)
		fmt.Println("Success login: ", username, password)
	} else {
		fmt.Println("Failed login")
	}
}

func TestSQLParameterInsert(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := SQLParameterGetConnection(driver, source)
	defer db.Close()

	ctx := context.Background()
	username := "kusuma'; DROP TABLE user; #"
	password := "wardhana"
	sqlQuery := `insert into user (username, password) values(?, ?)`

	if _, err := db.ExecContext(ctx, sqlQuery, username, password); err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data")
}

func TestSQLParameterUpdate(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := SQLParameterGetConnection(driver, source)
	defer db.Close()

	username := "noval"
	newUsername := "novalwardhana"
	newPassword := "secret12568"
	ctx := context.Background()
	sqlQuery := "update user set username = ?, password = ? where username = ?"
	if _, err := db.ExecContext(ctx, sqlQuery, newUsername, newPassword, username); err != nil {
		panic(err)
	}

	fmt.Println("Success update data")
}
