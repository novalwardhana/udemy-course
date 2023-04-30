package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
