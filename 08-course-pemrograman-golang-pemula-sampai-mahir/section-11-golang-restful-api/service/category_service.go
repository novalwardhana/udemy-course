package service

import (
	"belajar-golang-restful-api/model/web"
	"context"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryID int)
	FindByID(ctx context.Context, categoryID int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
