package server

import (
	"net/http"
	"restful-api/middleware"
)

func NewCategoryServer(middleware *middleware.CategoryMiddleware) *http.Server {
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware,
	}
	return &server
}
