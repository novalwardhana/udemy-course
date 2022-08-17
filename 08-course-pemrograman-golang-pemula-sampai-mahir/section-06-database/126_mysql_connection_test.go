package belajar_golang_database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestMysqlConnection(t *testing.T) {
	db, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_one")
	defer db.Close()
	if err != nil {
		panic(err)
	}
}

func TestMysqlMultiConnection(t *testing.T) {

	dbOne, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_one")
	defer dbOne.Close()
	if err != nil {
		message := fmt.Sprintf("Error connect to database test_one: %s", err.Error())
		fmt.Println(message)
		panic(err)
	}

	dbTwo, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_two")
	defer dbTwo.Close()
	if err != nil {
		message := fmt.Sprintf("Error connect to database test_two: %s", err.Error())
		fmt.Println(message)
		panic(err)
	}
}
