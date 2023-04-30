package simple2

import "errors"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

type UserService struct {
	UserRepository *UserRepository
	Error          bool
}

func NewUserService(userRepository *UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
		Error:          true,
	}
}

type UserController struct {
	UserService *UserService
}

func NewUserController(userService *UserService) (*UserController, error) {
	if userService.Error {
		return nil, errors.New("Test user")
	}

	return &UserController{
		UserService: userService,
	}, nil
}
