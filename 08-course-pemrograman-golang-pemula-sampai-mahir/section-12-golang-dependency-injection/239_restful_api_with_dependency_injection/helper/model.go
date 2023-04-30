package helper

import (
	"restful-api/model/domain"
	"restful-api/model/web"
)

func ToCategoryResponses(datas []domain.Category) []web.Category {
	var categories []web.Category
	for _, data := range datas {
		categories = append(categories, web.Category{
			ID:   data.ID,
			Name: data.Name,
		})
	}

	return categories
}
