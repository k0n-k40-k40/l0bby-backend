package handlers

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	. "l0bby_backend/internal/controllers/user"
	. "l0bby_backend/internal/models/user"
	. "l0bby_backend/internal/utils"

	_ "github.com/go-sql-driver/mysql"
)

const (
	databaseSource = "root:l0bby@tcp(172.17.0.2:3306)/l0bby"
)

func UserRegister(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", databaseSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusInternalServerError)
		fmt.Println(err)

		return
	}

	data, err := ParseJsonBody(body)
	if err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusInternalServerError)
		fmt.Println(err)
	}

	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Email:    data["email"].(string),
	}

	err = CreateUser(db, &user)
	if err != nil {
		http.Error(writer, "Failed to create user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "User \""+user.Username+"\" created")
}

func UserLogin(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", databaseSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusInternalServerError)
		fmt.Println(err)
	}

	data, err := ParseJsonBody(body)
	if err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusInternalServerError)
		fmt.Println(err)
	}

	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}

	err = LoginUser(db, &user)
	if err != nil {
		http.Error(writer, "Failed to login user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "User \""+user.Username+"\" logged in")
}
