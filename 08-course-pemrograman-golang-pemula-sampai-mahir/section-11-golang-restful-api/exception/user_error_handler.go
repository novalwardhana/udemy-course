package exception

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"net/http"
)

func UserErrorHandler(writer http.ResponseWriter, request *http.Request, error interface{}) {

	if checkUserNotFoundError(writer, request, error) {
		return
	}

	if checkUserValidationError(writer, request, error) {
		return
	}

	writer.WriteHeader(http.StatusInternalServerError)
	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
	}
	helper.UserWriteResponseBody(writer, webResponse)
}

func checkUserNotFoundError(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	userNotFoundError, ok := error.(UserNotFoundError)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   userNotFoundError,
		}
		writer.WriteHeader(http.StatusNotFound)
		helper.UserWriteResponseBody(writer, webResponse)
	}

	return ok
}

func checkUserValidationError(writer http.ResponseWriter, request *http.Request, error interface{}) bool {
	userValidationError, ok := error.(UserValidationError)
	if ok {
		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   userValidationError,
		}
		writer.WriteHeader(http.StatusBadRequest)
		helper.UserWriteResponseBody(writer, webResponse)
	}

	return ok
}
