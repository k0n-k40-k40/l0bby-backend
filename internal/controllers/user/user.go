package user

import (
	"database/sql"

	. "l0bby_backend/internal/models/user"

	_ "github.com/go-sql-driver/mysql"
)

func CreateUser(db *sql.DB, user *User) error {
	query := "INSERT INTO users (username, password, email) VALUES (?, ?, ?)"
	_, err := db.Exec(query, user.Username, user.Password, user.Email)
	return err
}

func LoginUser(db *sql.DB, user *User) error {
	query := "SELECT * FROM users WHERE username = ? AND password = ?"
	err := db.QueryRow(query, user.Username, user.Password).Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return err
}
