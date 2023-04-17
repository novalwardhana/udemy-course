package service

import (
	"belajar-golang-restful-api/model/web"
	"context"
)

type UserService interface {
	GetAllData(ctx context.Context) []web.UserResponse
	GetByID(ctx context.Context, id int) web.UserResponse
	Create(ctx context.Context, params web.UserCreateRequest) web.UserResponse
	Update(ctx context.Context, params web.UserUpdateRequest) web.UserResponse
	Delete(ctx context.Context, id int)
}
