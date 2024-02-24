package user

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	utils "l0bby_backend/internal/utils"

	_ "github.com/go-sql-driver/mysql"
)

func HandleRegister(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", utils.DatabaseSource)
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

	data, err := utils.ParseJsonBody(body)
	if err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusInternalServerError)
		fmt.Println(err)
	}

	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
		Email:    data["email"].(string),
	}

	err = createUser(db, &user)
	if err != nil {
		http.Error(writer, "Failed to create user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "User \""+user.Username+"\" created")
}

func HandleLogin(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", utils.DatabaseSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	body, err := io.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "Failed to read request body", http.StatusInternalServerError)
		fmt.Println(err)
	}

	data, err := utils.ParseJsonBody(body)
	if err != nil {
		http.Error(writer, "Failed to parse request body", http.StatusInternalServerError)
		fmt.Println(err)
	}

	user := User{
		Username: data["username"].(string),
		Password: data["password"].(string),
	}

	err = loginUser(db, &user)
	if err != nil {
		http.Error(writer, "Failed to login user", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "User \""+user.Username+"\" logged in")
}
