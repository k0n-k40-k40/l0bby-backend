package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"l0bby_backend/internal/user"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/register", user.HandleRegister)
	router.Post("/login", user.HandleLogin)

	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", router)
}
