package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryConteroller(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.CategoryReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.CategoryWriterToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.CategoryReadFromRequestBody(request, &categoryUpdateRequest)

	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helper.PanicIfError(err)
	categoryUpdateRequest.ID = categoryID

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.CategoryWriterToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), categoryID)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.CategoryWriterToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindByID(request.Context(), categoryID)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	helper.CategoryWriterToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}

	helper.CategoryWriterToResponseBody(writer, webResponse)
}
