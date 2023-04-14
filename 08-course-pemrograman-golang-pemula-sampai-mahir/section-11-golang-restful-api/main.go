package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validator := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validator)
	categoryController := controller.NewCategoryConteroller(categoryService)

	router := httprouter.New()
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryID", categoryController.FindByID)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryID", categoryController.Update)
	router.DELETE("/api/categories/:categoryID", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
