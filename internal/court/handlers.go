package court

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	. "l0bby_backend/internal/sportstype"
	utils "l0bby_backend/internal/utils"

	_ "github.com/go-sql-driver/mysql"
)

func HandleCreateCourt(writer http.ResponseWriter, request *http.Request) {
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

	court := Court{
		Name:    data["name"].(string),
		Type:    SportsType(data["type"].(float64)),
		Address: data["address"].(string),
		Area:    data["area"].(string),
		Phone:   data["phone"].(string),
	}

	err = createCourt(db, &court)
	if err != nil {
		http.Error(writer, "Failed to create court", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	fmt.Fprintf(writer, "Court \""+court.Name+"\" created")
}

func HandleGetAllCourts(writer http.ResponseWriter, request *http.Request) {
	db, err := sql.Open("mysql", utils.DatabaseSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	courts, err := getAllCourts(db)
	if err != nil {
		http.Error(writer, "Failed to get courts", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	marshalledCourts, err := json.Marshal(courts)
	if err != nil {
		http.Error(writer, "Failed to marshal courts", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write(marshalledCourts)
}
