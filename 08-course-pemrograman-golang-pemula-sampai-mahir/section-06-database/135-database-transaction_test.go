package belajar_golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func DatabaseTransactionCreateConnection(driver, source string) *sql.DB {
	db, err := sql.Open(driver, source)
	if err != nil {
		panic(err)
	}
	return db
}

func TestDatabaseTransaction(t *testing.T) {
	driver := "mysql"
	source := "noval:noval@tcp(localhost:3306)/test_one"
	db := DatabaseTransactionCreateConnection(driver, source)
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	var IDs []int64
	prepare, err := tx.PrepareContext(ctx, "insert into comments (email, comment) values(?, ?) ")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 100; i++ {
		email := fmt.Sprintf("test_%s@gmail.com", time.Now().Format("20060102"))
		comment := fmt.Sprintf("Comment - %s", time.Now().Format("20060102"))

		result, err := prepare.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}
		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}
		IDs = append(IDs, id)

		time.Sleep(1 * time.Second)
	}

	fmt.Println("New transaction ID: ", IDs)
	// err = tx.Commit()
	// if err != nil {
	// 	panic(err)
	// }
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}

}
