package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
)

func UpdateUsername(u *structures.UpdateUsername) error {
	db, err := sql.Open("sqlite3", "database/db.db")
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE User SET username=? WHERE id_user=?", u.NewUsername, u.IdUser)
	if err != nil {
		return err
	}
	return nil
}
