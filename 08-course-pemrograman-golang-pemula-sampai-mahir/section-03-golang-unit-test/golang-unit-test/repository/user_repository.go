package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type UserMock struct {
	Mock mock.Mock
}

type User interface {
	GetUser(id int) *entity.User
}

func (u *UserMock) GetUser(id int) *entity.User {
	result := u.Mock.Called(id)
	if result.Get(0) == nil {
		return nil
	}
	return result.Get(0).(*entity.User)
}
