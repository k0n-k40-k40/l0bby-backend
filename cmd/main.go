package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"l0bby_backend/internal/controllers/handlers"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", handlers.UserRegister)
	router.Post("/login", handlers.UserLogin)

	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", router)
}
