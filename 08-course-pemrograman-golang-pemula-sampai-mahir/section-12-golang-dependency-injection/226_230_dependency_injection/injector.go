//go:build wireinject
// +build wireinject

package simple2

import "github.com/google/wire"

func InitializeUser() (*UserController, error) {
	wire.Build(NewUserRepository, NewUserService, NewUserController)
	return nil, nil
}

func InitializeCategory() (CategoryController, error) {
	wire.Build(NewCategoryService, NewCategoryController, NewCategoryRepository)
	return nil, nil
}
