package repository

import (
	"belajar-golang-database/136-repository-pattern/entity"
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func getConnection() *sql.DB {
	db, err := sql.Open("mysql", "noval:noval@tcp(localhost:3306)/test_one")
	if err != nil {
		panic(err)
	}
	return db
}

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(getConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "testnov@gmail.com",
		Comment: "Test email",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindbyID(t *testing.T) {
	commentRepository := NewCommentRepository(getConnection())

	ctx := context.Background()
	result, err := commentRepository.FindByID(ctx, 11144)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(getConnection())

	ctx := context.Background()
	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
