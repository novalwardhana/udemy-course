package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindByID(ctx context.Context, tx *sql.Tx, categoryID int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
