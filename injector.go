//go:build wireinject
// +build wireinject

package main

import (
	"belajar-go-rest-api/app"
	"belajar-go-rest-api/controller"
	"belajar-go-rest-api/middleware"
	"belajar-go-rest-api/repository"
	"belajar-go-rest-api/services"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	services.NewCategoryService,
	wire.Bind(new(services.CategoryService), new(*services.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitiaLizedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
