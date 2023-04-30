package service

import "restful-api/model/web"

type CategoryService interface {
	GetAll() []web.Category
}
