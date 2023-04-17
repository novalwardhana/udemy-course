package app

import (
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UserNewRouter(controller controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", controller.GetAllData)
	router.GET("/api/users/:id", controller.GetByID)
	router.POST("/api/users", controller.Create)
	router.PUT("/api/users/:id", controller.Update)
	router.DELETE("/api/users/:id", controller.Delete)
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		exception.UserErrorHandler(writer, request, error)
	}

	return router
}
