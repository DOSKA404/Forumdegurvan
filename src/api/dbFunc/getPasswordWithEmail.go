package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func GetPasswordWithEmail(email string) error {
	db, err := sql.Open("sqlite3", "database/db.db")
	tools.HandlerError(err)
	defer db.Close()
	return nil
}
