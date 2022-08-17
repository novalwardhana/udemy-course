package repository

import (
	"golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

type CategoryRepository interface {
	FindByID(id string) *entity.Category
}

func (repository *CategoryRepositoryMock) FindByID(id string) *entity.Category {
	result := repository.Mock.Called(id)
	if result.Get(0) == nil {
		return nil
	}
	category := result.Get(0).(entity.Category)
	return &category
}
