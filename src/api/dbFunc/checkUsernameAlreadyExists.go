package dbFunc

import (
	"database/sql"
	"errors"

	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func CheckUsernameAlreadyExists(db *sql.DB, username string) error {
	row := db.QueryRow("SELECT COUNT(1) FROM User WHERE username=?", username)
	var result int
	err := row.Scan(&result)
	if err != nil {
		tools.HandlerError(err)
	}
	if result != 0 {
		err := errors.New("username already exists")
		return err
	}
	return nil
}
