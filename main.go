package main

import (
	"learn-go-restful-api/app"
	"learn-go-restful-api/controller"
	"learn-go-restful-api/helper"
	"learn-go-restful-api/middleware"
	"learn-go-restful-api/repository"
	"learn-go-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	authMiddleware := middleware.NewAuthMiddleWare(router)

	server := http.Server{
		Addr:    "localhost:9090",
		Handler: authMiddleware,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
