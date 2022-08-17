package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

/* Table
CREATE TABLE comments
(
	id INT NOT NULL AUTO_INCREMENT,
	email VARCHAR(100) NOT NULL,
	comment TEXT,
	PRIMARY KEY (id)
)
*/

func AutoIncreamentGetConnection(driver, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)

	return db
}

func TestAutoIncreamentSingleData(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := AutoIncreamentGetConnection(driver, source)
	defer db.Close()

	ctx := context.Background()
	sampleEmail := "root@localhost.dev"
	sampleComment := "DROP TABLE comments; # ' ' ' \n "
	sqlQuery := `insert into comments (email, comment) values(?, ?)`

	result, err := db.ExecContext(ctx, sqlQuery, sampleEmail, sampleComment)
	if err != nil {
		panic(err)
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Last ID: ", lastID)
}

func TestAutoIncrementMultiData(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := AutoIncreamentGetConnection(driver, source)
	defer db.Close()

	type Data struct {
		Email   string
		Comment string
	}
	var datas = []Data{
		{"noval@gmail.com", "DROP TABLE customer; # ''"},
		{"kusuma@gmail.com", "Test comment 2"},
		{"wardhana@gmail.com", "Test comment 3"},
	}
	ctx := context.Background()

	var ID []int64
	sqlQuery := "insert into comments (email, comment) values(?, ?)"
	for _, data := range datas {
		result, err := db.ExecContext(ctx, sqlQuery, data.Email, data.Comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		ID = append(ID, id)
	}

	fmt.Println("Success insert new data: ", ID)
}
