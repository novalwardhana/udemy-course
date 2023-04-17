package helper

import (
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		FullName: user.FullName,
	}
}

func ToUserResponses(users []domain.User) []web.UserResponse {
	var userResponses []web.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, web.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			Username: user.Username,
			FullName: user.FullName,
		})
	}
	return userResponses
}
