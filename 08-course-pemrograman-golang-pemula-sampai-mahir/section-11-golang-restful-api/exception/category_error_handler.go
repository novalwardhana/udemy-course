package exception

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func CategoryErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {

	if checkCategoryNotFoundError(writer, request, err) {
		return
	}

	if checkCategoryvalidationErrors(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func checkCategoryNotFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(CategoryNotFoundError)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.Error,
		}
		helper.CategoryWriterToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func checkCategoryvalidationErrors(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		helper.CategoryWriterToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "Internal server error",
		Data:   err,
	}
	helper.CategoryWriterToResponseBody(writer, webResponse)
}
