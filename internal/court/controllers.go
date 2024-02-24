package court

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func createCourt(db *sql.DB, court *Court) error {
	query := "INSERT INTO courts (name, type, address, area, phone) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, court.Name, court.Type, court.Address, court.Area, court.Phone)
	return err
}

func getAllCourts(db *sql.DB) ([]Court, error) {
	query := "SELECT * FROM courts"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	courts := []Court{}
	for rows.Next() {
		var court Court
		err := rows.Scan(&court.ID, &court.Name, &court.Type, &court.Address, &court.Area, &court.Phone)
		if err != nil {
			return nil, err
		}
		courts = append(courts, court)
	}

	return courts, nil
}
