package app

import (
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CategoryNewRouter(categoryController controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryID", categoryController.FindByID)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryID", categoryController.Update)
	router.DELETE("/api/categories/:categoryID", categoryController.Delete)

	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		exception.CategoryErrorHandler(writer, request, err)
	}

	return router
}
