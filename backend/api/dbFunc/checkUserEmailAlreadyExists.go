package dbFunc

import (
	"database/sql"
	"errors"
)

func CheckUserEmailAlreadyExists(db *sql.DB, email string) error {
	row := db.QueryRow("SELECT COUNT(1) FROM User WHERE email=?", email)
	var result int
	err := row.Scan(&result)
	if err != nil {
		return err
	}
	if result != 0 {
		err := errors.New("user already exists")
		return err
	}
	return nil
}
