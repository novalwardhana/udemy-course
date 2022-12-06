package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
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

	ctx := context.Background()
	sqlQuery := "insert into comments (email, comment) values(?, ?)"
	stmt, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	IDs := []int64{}
	for i := 1; i <= 10; i++ {
		email := fmt.Sprintf("noval-%d@gmail.com", i)
		comment := fmt.Sprintf("test comment ke: %d", i)
		result, err := stmt.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		lastInsertID, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		IDs = append(IDs, lastInsertID)
	}
	fmt.Println("New comment ID: ", IDs)
}

func TestPrepareStatementBulkInsert(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := PrepareStatementGetConnection(driver, source)
	defer db.Close()

	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	ctx := context.Background()
	iteration := 10000

	sqlQuery := `insert into comments (email, comment) values(?, ?)`
	stmt, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}

	IDs := []int64{}
	for i := 1; i <= iteration; i++ {
		wg.Add(1)
		loop := i
		go func() {
			defer wg.Done()
			email := fmt.Sprintf("novalwardhana-%d@gmail.com", loop)
			comment := fmt.Sprintf("Test comment ke: %d", loop)

			mutex.Lock()
			result, err := stmt.ExecContext(ctx, email, comment)
			if err != nil {
				panic(err)
			}
			lastID, err := result.LastInsertId()
			if err != nil {
				panic(err)
			}
			IDs = append(IDs, lastID)
			mutex.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("New IDs: ", IDs)
}
