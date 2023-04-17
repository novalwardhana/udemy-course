package middleware

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type UserMiddleware struct {
	Handler http.Handler
}

func (u *UserMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	authHeader := request.Header.Get("Authorization")
	authHeaderArr := strings.Split(authHeader, " ")
	if len(authHeaderArr) != 2 {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		writer.WriteHeader(http.StatusUnauthorized)
		helper.UserWriteResponseBody(writer, webResponse)
		return
	}

	authDecode, err := base64.StdEncoding.DecodeString(authHeaderArr[1])
	if err != nil {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		writer.WriteHeader(http.StatusUnauthorized)
		helper.UserWriteResponseBody(writer, webResponse)
		return
	}

	authDecodeArr := strings.Split(string(authDecode), ":")
	if len(authDecodeArr) != 2 {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		writer.WriteHeader(http.StatusUnauthorized)
		helper.UserWriteResponseBody(writer, webResponse)
		return
	}

	if authDecodeArr[0] != "noval" || authDecodeArr[1] != "secret123" {
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		writer.WriteHeader(http.StatusUnauthorized)
		helper.UserWriteResponseBody(writer, webResponse)
		return
	}

	u.Handler.ServeHTTP(writer, request)
}

func NewUserMiddleware(router *httprouter.Router) *UserMiddleware {
	return &UserMiddleware{
		Handler: router,
	}
}
