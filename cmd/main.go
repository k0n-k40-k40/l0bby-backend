package main

import (
    _ "database/sql"
    _ "encoding/json"
    _ "fmt"
    "log"
    "net/http"
    _ "strconv"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

const (
    dbDriver = "mysql"
    dbUser   = "root"
    dbPass   = "l0bby"
    dbName   = "l0bby"
)

func main() {
    // Create a new router
    r := mux.NewRouter()

    // r.HandleFunc("/user", createUserHandler).Methods("POST")
    // r.HandleFunc("/user/{id}", getUserHandler).Methods("GET")
    // r.HandleFunc("/user/{id}", updateUserHandler).Methods("PUT")
    // r.HandleFunc("/user/{id}", deleteUserHandler).Methods("DELETE")

    log.Println("Server listening on :8090")
    log.Fatal(http.ListenAndServe(":8090", r))
}