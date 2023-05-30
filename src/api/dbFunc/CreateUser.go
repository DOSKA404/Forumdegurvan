package dbFunc

import (
	"database/sql"
	"encoding/hex"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func PutUserInDb(user *structures.User) error {
	db, err := sql.Open("sqlite3", "database/db.db")
	tools.HandlerError(err)
	defer db.Close()
	err = CheckUserEmailAlreadyExists(db, user.Email)
	if err != nil {
		return err
	}
	err = CheckUsernameAlreadyExists(db, user.Username)
	if err != nil {
		return err
	}
	records := "INSERT INTO User(email, username, date_of_birth, password_hash) VALUES(?, ?, ?, ?)"
	querry, err := db.Prepare(records)
	tools.HandlerError(err)
	_, err = querry.Exec(user.Email, user.Username, user.DateOfBirth, hex.EncodeToString(tools.HashPassword(user.Password)))
	tools.HandlerError(err)
	return nil
}
