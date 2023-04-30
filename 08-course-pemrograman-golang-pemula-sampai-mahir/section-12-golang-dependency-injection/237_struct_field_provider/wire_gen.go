// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package simple

// Injectors from injector.go:

func InitializeCategoryRepository() *CategoryRepository {
	categoryService := NewCategoryService()
	categoryRepository := categoryService.CategoryRepository
	return categoryRepository
}

func mu() *ManchesterUnited {
	club := NewClub()
	manchesterUnited := club.ManchesterUnited
	return manchesterUnited
}

func chelsea() *Chelsea {
	club := NewClub()
	simpleChelsea := club.Chelsea
	return simpleChelsea
}

func arsenal() *Arsenal {
	club := NewClub()
	simpleArsenal := club.Arsenal
	return simpleArsenal
}
