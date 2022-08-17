package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository = repository.UserMock{Mock: mock.Mock{}}
var userService = User{repository: &userRepository}

func TestUserSuccess(t *testing.T) {
	userRepository.Mock.On("GetUser", 1).Return(&entity.User{
		ID:       1,
		Email:    "novalwardhana@gmail.com",
		Password: "secret123",
		Name:     "Novalita Kusuma Wardhana",
	})
	result, err := userService.GetUser(1)
	assert.NotNil(t, result)
	assert.Nil(t, err)
}

func TestUserNil(t *testing.T) {
	userRepository.Mock.On("GetUser", 2).Return(nil)
	result, err := userService.GetUser(2)
	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestUserTable(t *testing.T) {
	users := []*entity.User{
		&entity.User{
			ID: 1, Email: "noval@gmail.com", Password: "Test", Name: "Noval",
		},
		&entity.User{
			ID: 2, Email: "kusuma@gmail.com", Password: "Test", Name: "Kusuma",
		},
		&entity.User{
			ID: 3, Email: "wardhana@gmail.com", Password: "Test", Name: "Wardhana",
		},
	}

	for _, user := range users {
		t.Run(user.Name, func(t *testing.T) {
			userRepository.Mock.On("GetUser", user.ID).Return(user)
			result, err := userService.GetUser(user.ID)
			assert.NotNil(t, result)
			assert.Nil(t, err)
		})
	}
}
