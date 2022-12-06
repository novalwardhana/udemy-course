package repository

import (
	"belajar-golang-database/136-repository-pattern/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type commentRepository struct {
	DB *sql.DB
}

type CommentRepository interface {
	Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error)
	FindByID(ctx context.Context, id int32) (entity.Comment, error)
	FindAll(ctx context.Context) ([]entity.Comment, error)
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepository{
		DB: db,
	}
}

func (repo *commentRepository) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sql := "insert into comments (email, comment) values(?, ?) "
	result, err := repo.DB.ExecContext(ctx, sql, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.ID = int32(id)

	return comment, nil
}

func (repo *commentRepository) FindByID(ctx context.Context, id int32) (entity.Comment, error) {

	var comment entity.Comment
	sql := "select * from comments where id = ? limit 1"
	rows, err := repo.DB.QueryContext(ctx, sql, id)
	if err != nil {
		return comment, err
	}

	defer rows.Close()
	if rows.Next() {
		err := rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		if err != nil {
			return comment, err
		}
		return comment, nil
	} else {
		return comment, errors.New(fmt.Sprintf("ID: %d, not found", id))
	}
}

func (repo *commentRepository) FindAll(ctx context.Context) ([]entity.Comment, error) {
	var comments []entity.Comment
	sql := "select * from comments"
	rows, err := repo.DB.QueryContext(ctx, sql)
	if err != nil {
		return comments, err
	}

	for rows.Next() {
		var comment entity.Comment
		err := rows.Scan(&comment.ID, &comment.Email, &comment.Comment)
		if err != nil {
			return comments, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}
