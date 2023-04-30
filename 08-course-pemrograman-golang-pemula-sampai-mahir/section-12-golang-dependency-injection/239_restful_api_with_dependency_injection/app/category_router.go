package app

import (
	"restful-api/controller"

	"github.com/julienschmidt/httprouter"
)

func CategoryRouter(categoryController controller.CategoryController) *httprouter.Router {

	router := httprouter.New()
	router.GET("/api/category", categoryController.GetAll)

	return router
}
