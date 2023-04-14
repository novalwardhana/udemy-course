package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	sql := "insert into category (name) values(?)"
	result, err := tx.ExecContext(ctx, sql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.ID = int(id)

	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	sql := "update category set name = ? where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.ID)
	helper.PanicIfError(err)

	return category

}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {

	sql := "delete from category where id = ?"
	_, err := tx.ExecContext(ctx, sql, category.ID)
	helper.PanicIfError(err)

}

func (repository *CategoryRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error) {

	sql := "select id, name from category where id = ?"
	rows, err := tx.QueryContext(ctx, sql, categoryID)
	helper.PanicIfError(err)

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}

}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {

	sql := "select id, name from category"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(
			&category.ID,
			&category.Name,
		)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
