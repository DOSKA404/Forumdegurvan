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
	rows := db.QueryRow("SELECT id_user, email, username, date_of_birth FROM User WHERE email=?", email)
	tools.HandlerError(err)
	var user structures.User
	err = rows.Scan(&user.Id_user, &user.Email, &user.Username, &user.DateOfBirth)
	tools.HandlerError(err)
	return user
}
