package belajar_golang_database

import (
	"database/sql"
	"fmt"
	"sync"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DatabasePoolingGetConnection() *sql.DB {
	db, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_one")
	if err != nil {
		fmt.Println("Ada error: ", err.Error())
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	// db.SetConnMaxIdleTime(2 * time.Minute) Not exist
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}

func TestDatabasePooling(t *testing.T) {
	dbConn := DatabasePoolingGetConnection()
	defer dbConn.Close()
}

func DatabasePoolingParamConnection(driver string, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(5 * time.Minute)
	return db
}

func TestDatabaseMultiplePooling(t *testing.T) {
	type Database struct {
		driver string
		source string
	}
	databases := []Database{
		{"mysql", "noval:noval@tcp(localhost:3306)/test_one"},
		{"mysql", "noval:noval@tcp(localhost:3306)/test_two"},
	}

	wg := &sync.WaitGroup{}
	for _, data := range databases {
		wg.Add(1)
		database := data
		go func() {
			defer wg.Done()
			db := DatabasePoolingParamConnection(database.driver, database.source)
			time.Sleep(10 * time.Second)

			db.Close()
		}()
	}
	wg.Wait()

	fmt.Println("Success test db connection")
}
