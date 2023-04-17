package service

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator"
)

type UserServiceImpl struct {
	db       *sql.DB
	repo     repository.UserRepository
	validate *validator.Validate
}

func NewUserService(db *sql.DB, repo repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		db:       db,
		repo:     repo,
		validate: validate,
	}
}

func (service *UserServiceImpl) GetAllData(ctx context.Context) []web.UserResponse {

	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	users := service.repo.GetAllData(ctx, tx)
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	userResponses := helper.ToUserResponses(users)
	return userResponses
}

func (service *UserServiceImpl) GetByID(ctx context.Context, id int) web.UserResponse {

	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	user, err := service.repo.GetByID(ctx, tx, id)
	if err != nil {
		panic(err)
	}

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	userResponse := helper.ToUserResponse(user)
	return userResponse
}

func (service *UserServiceImpl) Create(ctx context.Context, params web.UserCreateRequest) web.UserResponse {

	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	err = service.validate.Struct(params)
	if err != nil {
		panic(err)
	}
	user := domain.User{
		Email:    params.Email,
		Username: params.Username,
		Password: params.Password,
		FullName: params.FullName,
	}
	user = service.repo.Create(ctx, tx, user)

	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	userResponse := helper.ToUserResponse(user)
	return userResponse

}

func (service *UserServiceImpl) Update(ctx context.Context, params web.UserUpdateRequest) web.UserResponse {

	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	err = service.validate.Struct(params)
	if err != nil {
		panic(err)
	}

	user, err := service.repo.GetByID(ctx, tx, params.ID)
	if err != nil {
		panic(err)
	}
	user.Email = params.Email
	user.Username = params.Username
	user.Password = params.Password
	user.FullName = params.FullName

	service.repo.Update(ctx, tx, user)
	err = tx.Commit()
	if err != nil {
		panic(err)
	}

	userResponse := helper.ToUserResponse(user)
	return userResponse
}

func (service *UserServiceImpl) Delete(ctx context.Context, id int) {
	tx, err := service.db.Begin()
	if err != nil {
		panic(err)
	}

	_, err = service.repo.GetByID(ctx, tx, id)
	if err != nil {
		panic(err)
	}

	service.repo.Delete(ctx, tx, id)
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
