package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	_ "github.com/mattn/go-sqlite3"
)

func GetUserInfo(email string) (structures.User, error) {
	db, err := sql.Open("sqlite3", "database/db.sqlite")
	if err != nil {
		return structures.User{}, err
	}
	defer db.Close()
	rows := db.QueryRow("SELECT id_user, email, username, date_of_birth FROM User WHERE email=?", email)
	if err != nil {
		return structures.User{}, err
	}
	var user structures.User
	err = rows.Scan(&user.Id_user, &user.Email, &user.Username, &user.DateOfBirth)
	if err != nil {
		return structures.User{}, err
	}
	return user, nil
}
