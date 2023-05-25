package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func PutUserInDb(user *structures.User) {
	db, err := sql.Open("sqlite3", "database/db.db")
	tools.HandlerError(err)
	defer db.Close()
	tools.HandlerError(err)
	records := "INSERT INTO User(email, username, date_of_birth, password_hash) VALUES(?, ?, ?, ?)"
	querry, err := db.Prepare(records)
	tools.HandlerError(err)
	_, err = querry.Exec(user.Email, user.Username, user.DateOfBirth, user.Password)
	tools.HandlerError(err)
}
