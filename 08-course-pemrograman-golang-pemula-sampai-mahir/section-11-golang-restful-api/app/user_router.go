package app

import (
	"belajar-golang-restful-api/controller"

	"github.com/julienschmidt/httprouter"
)

func UserNewRouter(controller controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/users", controller.GetAllData)
	router.GET("/api/users/:id", controller.GetByID)
	router.POST("/api/users", controller.Create)
	router.PUT("/api/users/:id", controller.Update)
	router.DELETE("/api/users/:id", controller.Delete)

	return router
}
