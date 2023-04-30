package simple2

import (
	"errors"
)

type CategoryRepository interface {
}

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

type CategoryService interface {
}

type CategoryServiceImpl struct {
	CategoryRepository CategoryRepository
	Error              bool
}

func NewCategoryService(categoryRepository CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		Error:              false,
	}
}

type CategoryController interface {
}

type CategoryControllerImpl struct {
	CategoryService CategoryService
}

func NewCategoryController(categoryService CategoryService) (CategoryController, error) {
	categoryServiceImpl := categoryService.(*CategoryServiceImpl)
	if categoryServiceImpl.Error {
		return nil, errors.New("Error category")
	}
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}, nil
}
