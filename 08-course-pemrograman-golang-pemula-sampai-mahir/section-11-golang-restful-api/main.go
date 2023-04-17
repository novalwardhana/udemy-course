package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/middleware"
	"belajar-golang-restful-api/repository"
	"belajar-golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func main() {

	forever := make(chan int)
	go func() {

		/* HTTP 1 */
		go func() {
			db := app.NewDB()
			validator := validator.New()
			categoryRepository := repository.NewCategoryRepository()
			categoryService := service.NewCategoryService(categoryRepository, db, validator)
			categoryController := controller.NewCategoryConteroller(categoryService)
			router := app.CategoryNewRouter(categoryController)

			server := http.Server{
				Addr:    "localhost:3000",
				Handler: middleware.NewCategoryAuthMiddleware(router),
			}
			err := server.ListenAndServe()
			helper.PanicIfError(err)
		}()

		/* HTTP 2 */
		go func() {
			db := app.NewDB()
			validator := validator.New()
			userRepository := repository.NewUserRepository()
			userService := service.NewUserService(db, userRepository, validator)
			userController := controller.NewUserController(*validator, userService)
			router := app.UserNewRouter(userController)

			server := http.Server{
				Addr:    "localhost:8080",
				Handler: router,
			}
			err := server.ListenAndServe()
			if err != nil {
				panic(err)
			}
		}()

	}()

	<-forever

}
