package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func GetUserInfo(email string) structures.User {
	db, err := sql.Open("sqlite3", "database/db.db")
	tools.HandlerError(err)
	defer db.Close()
	rows := db.QueryRow("SELECT User.Email, User.Username, User.DateOfBirth FROM User WHERE email=?", email)
	tools.HandlerError(err)
	var user structures.User
	err = rows.Scan(&user.Email, &user.Username, &user.DateOfBirth)
	tools.HandlerError(err)
	return user
}
