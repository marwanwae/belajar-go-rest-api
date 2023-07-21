package main

import (
	"belajar-go-rest-api/helper"
	"belajar-go-rest-api/middleware"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8081",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitiaLizedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
