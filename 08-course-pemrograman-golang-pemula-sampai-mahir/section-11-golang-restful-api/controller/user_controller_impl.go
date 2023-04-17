package controller

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/service"
	"context"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	validate validator.Validate
	service  service.UserService
}

func NewUserController(validate validator.Validate, service service.UserService) UserController {
	return &UserControllerImpl{
		validate: validate,
		service:  service,
	}
}

func (controller *UserControllerImpl) GetAllData(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	users := controller.service.GetAllData(context.Background())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   users,
	}
	helper.UserWriteResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) GetByID(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	user := controller.service.GetByID(context.Background(), userID)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	}
	helper.UserWriteResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var createRequest web.UserCreateRequest
	helper.UserReadFromRequestBody(request, &createRequest)

	user := controller.service.Create(context.Background(), createRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	}
	helper.UserWriteResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	var updateRequest web.UserUpdateRequest
	helper.UserReadFromRequestBody(request, &updateRequest)
	updateRequest.ID = userID

	user := controller.service.Update(context.Background(), updateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   user,
	}
	helper.UserWriteResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	controller.service.Delete(context.Background(), userID)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	helper.UserWriteResponseBody(writer, webResponse)
}
