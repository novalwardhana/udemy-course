package service

import (
	"context"
	"database/sql"
	"restful-api/helper"
	"restful-api/model/web"
	"restful-api/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	DB                 *sql.DB
	CategoryRepository repository.CategoryRepository
	Validate           *validator.Validate
}

func NewCategoryService(db *sql.DB, categoryRepository repository.CategoryRepository, validate *validator.Validate) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		DB:                 db,
		CategoryRepository: categoryRepository,
		Validate:           validate,
	}
}

func (c *CategoryServiceImpl) GetAll() []web.Category {
	tx, err := c.DB.Begin()
	helper.CategoryPanicIfError(err)
	defer tx.Commit()

	categories := c.CategoryRepository.GetAll(context.Background(), tx)
	return helper.ToCategoryResponses(categories)
}
