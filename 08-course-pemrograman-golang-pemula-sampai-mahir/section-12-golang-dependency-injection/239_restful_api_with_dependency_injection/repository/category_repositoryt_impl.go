package repository

import (
	"context"
	"database/sql"
	"restful-api/helper"
	"restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	sql := `select id, name from category`
	rows, err := tx.QueryContext(ctx, sql)
	helper.CategoryPanicIfError(err)
	defer rows.Close()

	var datas []domain.Category
	for rows.Next() {
		var data domain.Category
		err := rows.Scan(
			&data.ID,
			&data.Name,
		)
		helper.CategoryPanicIfError(err)
		datas = append(datas, data)
	}

	return datas
}
