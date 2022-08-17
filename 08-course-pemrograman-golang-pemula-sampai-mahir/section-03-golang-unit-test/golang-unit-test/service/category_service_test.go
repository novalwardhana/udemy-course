package service

import (
	"golang-unit-test/entity"
	"golang-unit-test/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceNotFound(t *testing.T) {
	categoryRepository.Mock.On("FindByID", "1").Return(nil)
	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServiceFound(t *testing.T) {
	params := entity.Category{
		ID:   "23",
		Name: "Noval Wardhana",
	}
	categoryRepository.Mock.On("FindByID", "23").Return(params)
	category, err := categoryService.Get("23")
	assert.NotNil(t, category)
	assert.Nil(t, err)
	assert.Equal(t, params.ID, category.ID)
	assert.Equal(t, params.Name, category.Name)
}
