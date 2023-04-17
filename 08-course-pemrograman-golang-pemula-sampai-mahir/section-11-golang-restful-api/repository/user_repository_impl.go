package repository

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repo *UserRepositoryImpl) GetAllData(ctx context.Context, tx *sql.Tx) []domain.User {
	sql := `select id, email, username, full_name from user order by id`
	rows, err := tx.QueryContext(ctx, sql)
	helper.UserPanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Username,
			&user.FullName,
		)
		helper.UserPanicIfError(err)
		users = append(users, user)
	}

	return users
}

func (repo *UserRepositoryImpl) GetByID(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	sql := `select id, email, username, full_name from user where id = ?`
	rows, err := tx.QueryContext(ctx, sql, id)
	helper.UserPanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Username,
			&user.FullName,
		)
		helper.UserPanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("User not found")
	}
}

func (repo *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := `insert into user (email, username, password, full_name) values(?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, sql, user.Email, user.Username, user.Password, user.FullName)
	helper.UserPanicIfError(err)

	id, err := result.LastInsertId()
	helper.UserPanicIfError(err)
	user.ID = int(id)

	return user
}

func (repo *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	sql := `update user set email = ?, username = ?, password = ?, full_name = ? where id = ?`
	_, err := tx.ExecContext(ctx, sql, user.Email, user.Username, user.Password, user.FullName, user.ID)
	helper.UserPanicIfError(err)

	return user
}

func (repo *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	sql := `delete from user where id = ?`
	_, err := tx.ExecContext(ctx, sql, id)
	helper.UserPanicIfError(err)
}
