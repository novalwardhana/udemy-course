package repository

import (
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	GetAllData(ctx context.Context, tx *sql.Tx) []domain.User
	GetByID(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, id int)
}
