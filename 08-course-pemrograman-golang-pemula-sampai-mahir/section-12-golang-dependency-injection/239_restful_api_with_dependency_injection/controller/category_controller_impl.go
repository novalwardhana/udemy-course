package controller

import (
	"net/http"
	"restful-api/model/web"
	"restful-api/service"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryControllerImpl {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (c *CategoryControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categories := c.CategoryService.GetAll()
	response := web.CategoryResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   categories,
	}

	web.CategoryWriterResponse(writer, response)
}
