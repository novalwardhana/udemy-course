package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func SQLExecutionGetConnection() *sql.DB {
	db, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_one")
	if err != nil {
		panic(err)
	}
	return db
}

func TestSQLExecution(t *testing.T) {
	db := SQLExecutionGetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := `insert into customer(id, name) values ('kusuma', 'wardhana')`
	_, err := db.ExecContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data")
}

func SQLExecutionParamConnection(driver string, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(2 * time.Minute)

	return db
}

func TestSQLMultipleExecution(t *testing.T) {
	dbDriver := "mysql"
	dbSource := "noval:noval@tcp(localhost:3306)/test_one"
	db := SQLExecutionParamConnection(dbDriver, dbSource)
	defer db.Close()

	names := []string{"Novalita", "Kusuma", "Wardhana"}
	for _, name := range names {
		id := strconv.Itoa(int(time.Now().Unix()))
		query := `insert into customer (id, name) values(?, ?)`

		ctx := context.Background()
		if _, err := db.ExecContext(ctx, query, id, name); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
