//go:build wireinject

// + wireinject
package run_wire

import (
	"net/http"
	"restful-api/app"
	"restful-api/controller"
	"restful-api/middleware"
	"restful-api/repository"
	"restful-api/server"
	"restful-api/service"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedCategoryServer() *http.Server {
	wire.Build(
		app.Database,
		validator.New,
		categorySet,
		app.CategoryRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewCategoryMiddleware,
		server.NewCategoryServer,
	)
	return nil
}
