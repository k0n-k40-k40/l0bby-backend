package court

import (
	"database/sql"
	. "l0bby_backend/internal/sportstype"

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

func getCourts_Area(db *sql.DB, area string) ([]Court, error) {
	query := "SELECT * FROM courts WHERE area = ?"
	rows, err := db.Query(query, area)

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

func getCourts_Type(db *sql.DB, sportsType string) ([]Court, error) {
	query := "SELECT * FROM courts WHERE type = ?"
	rows, err := db.Query(query, sportsType)

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

func getCourts_AreaType(db *sql.DB, area string, sportsType SportsType) ([]Court, error) {
	query := "SELECT * FROM courts WHERE area = ? AND type = ?"
	rows, err := db.Query(query, area, sportsType)

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
