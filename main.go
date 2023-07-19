package main

import (
	"belajar-go-rest-api/app"
	"belajar-go-rest-api/controller"
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/middleware"
	"belajar-go-rest-api/repository"
	"belajar-go-rest-api/services"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
