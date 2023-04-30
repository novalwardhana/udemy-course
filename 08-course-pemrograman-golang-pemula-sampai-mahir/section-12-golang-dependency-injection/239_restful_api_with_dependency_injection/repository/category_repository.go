package repository

import (
	"context"
	"database/sql"
	"restful-api/model/domain"
)

type CategoryRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
