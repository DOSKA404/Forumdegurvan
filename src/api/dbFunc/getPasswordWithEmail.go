package dbFunc

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func GetPasswordWithEmail(email string) (string, error) {
	result := ""
	db, err := sql.Open("sqlite3", "database/db.db")
	if err != nil {
		db.Close()
		return "", err
	}
	defer db.Close()
	row := db.QueryRow("SELECT password_hash FROM User WHERE email=?", email)
	row.Scan(&result)
	return result, nil
}
