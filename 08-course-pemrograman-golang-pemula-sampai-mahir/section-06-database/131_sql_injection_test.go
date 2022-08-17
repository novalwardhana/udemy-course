package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

/* Database sample
CREATE TABLE user
(
	username VARCHAR(100) NOT NULL,
	password VARCHAR(100) NOT NULL,
	PRIMARY KEY (username)
) ENGINE = InnoDB;
*/

func SQLInjectionGetConnection() *sql.DB {
	db, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_one")
	if err != nil {
		panic(err)
	}
	return db
}

func TestSQLInjectionExample(t *testing.T) {
	username := "noval'; #"
	password := "secret123"
	ctx := context.Background()

	db := SQLInjectionGetConnection()
	sqlQuery := "SELECT username FROM user WHERE username = '" + username +
		"' AND password = '" + password + "' LIMIT 1"
	fmt.Println("Query: ", sqlQuery)
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		rows.Scan(&username)
		fmt.Println("Sukses login:", username)
	} else {
		fmt.Println("Gagal login")
	}
}

func TestSQLInjectionByPassword(t *testing.T) {
	username := "noval"
	password := "secret123' # "
	db := SQLInjectionGetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "select username, password from user where password='" + password + "' and username ='" + username + "' LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	if rows.Next() {
		var username string
		var password string
		rows.Scan(&username, &password)
		fmt.Println("Success login: ", username, password)
	} else {
		fmt.Println("Failed login")
	}
}
