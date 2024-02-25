package party

import (
	"database/sql"
)

func createParty(db *sql.DB, party *Party) error {
	query := "INSERT INTO parties (name, code, court_id, start_time, end_time) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, party.Name, party.Code, party.Court.ID, party.StartTime, party.EndTime)
	return err
}

func getParty_Code(db *sql.DB, code string) (*Party, error) {
	query := "SELECT * FROM parties WHERE code = ?"
	row := db.QueryRow(query, code)

	var party Party
	err := row.Scan(&party.ID, &party.Name, &party.Code, &party.Court.ID, &party.StartTime, &party.EndTime)
	if err != nil {
		return nil, err
	}

	return &party, nil
}

func getParty_CourtArea(db *sql.DB, area string) ([]Party, error) {
	query := "SELECT * FROM parties JOIN courts ON parties.court_id = courts.id WHERE courts.area = ?"
	rows, err := db.Query(query, area)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	parties := []Party{}
	for rows.Next() {
		var party Party
		err := rows.Scan(&party.ID, &party.Name, &party.Code, &party.Court.ID, &party.StartTime, &party.EndTime)
		if err != nil {
			return nil, err
		}
		parties = append(parties, party)
	}

	return parties, nil
}

func getParty_CourtType(db *sql.DB, sportsType string) ([]Party, error) {
	query := "SELECT * FROM parties JOIN courts ON parties.court_id = courts.id WHERE courts.type = ?"
	rows, err := db.Query(query, sportsType)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	parties := []Party{}
	for rows.Next() {
		var party Party
		err := rows.Scan(&party.ID, &party.Name, &party.Code, &party.Court.ID, &party.StartTime, &party.EndTime)
		if err != nil {
			return nil, err
		}
		parties = append(parties, party)
	}

	return parties, nil
}

func getParty_CourtAreaType(db *sql.DB, area string, sportsType string) ([]Party, error) {
	query := "SELECT * FROM parties JOIN courts ON parties.court_id = courts.id WHERE courts.area = ? AND courts.type = ?"
	rows, err := db.Query(query, area, sportsType)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	parties := []Party{}
	for rows.Next() {
		var party Party
		err := rows.Scan(&party.ID, &party.Name, &party.Code, &party.Court.ID, &party.StartTime, &party.EndTime)
		if err != nil {
			return nil, err
		}
		parties = append(parties, party)
	}

	return parties, nil
}
