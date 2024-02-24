package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"

	"l0bby_backend/internal/court"
	"l0bby_backend/internal/user"
)

func main() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Post("/user/register", user.HandleRegister)
	router.Post("/user/login", user.HandleLogin)

	router.Post("/court/new", court.HandleCreateCourt)
	router.Get("/court/all", court.HandleGetAllCourts)

	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", router)
}
