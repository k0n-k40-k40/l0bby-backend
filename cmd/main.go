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
	router.Get("/court/area/{area}", court.HandleGetCourts_Area)
	router.Get("/court/type/{type}", court.HandleGetCourts_Type)
	router.Get("/court/area/{area}/type/{type}", court.HandleGetCourts_AreaType)

	fmt.Println("Server is running on port 8090")
	http.ListenAndServe(":8090", router)
}
