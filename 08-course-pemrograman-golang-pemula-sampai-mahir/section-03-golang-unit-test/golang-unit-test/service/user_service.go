package service

import (
	"errors"
	"golang-unit-test/entity"
	"golang-unit-test/repository"
)

type User struct {
	repository repository.User
}

func (u *User) GetUser(id int) (*entity.User, error) {
	result := u.repository.GetUser(id)
	if result == nil {
		return nil, errors.New("User not found")
	}
	return result, nil
}
