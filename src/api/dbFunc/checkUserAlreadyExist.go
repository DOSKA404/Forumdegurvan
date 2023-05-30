package dbFunc

import (
	"database/sql"
	"errors"

	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func CheckUserAlreadyExist(db *sql.DB, email string) error {
	row := db.QueryRow("SELECT COUNT(1) FROM User WHERE email=?", email)
	var result int
	err := row.Scan(&result)
	if err != nil {
		tools.HandlerError(err)
	}
	if result != 0 {
		err := errors.New("user already exists")
		return err
	}
	return nil
}
