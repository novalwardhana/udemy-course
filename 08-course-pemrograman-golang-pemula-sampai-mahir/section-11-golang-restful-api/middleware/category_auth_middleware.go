package middleware

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"net/http"
)

type CategoryAuthMiddleware struct {
	Handler http.Handler
}

func NewCategoryAuthMiddleware(handler http.Handler) *CategoryAuthMiddleware {
	return &CategoryAuthMiddleware{
		Handler: handler,
	}
}

func (middleware *CategoryAuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") != "RAHASIA" {

		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTORIZED",
		}
		helper.CategoryWriterToResponseBody(writer, webResponse)
		return
	}

	middleware.Handler.ServeHTTP(writer, request)
}
