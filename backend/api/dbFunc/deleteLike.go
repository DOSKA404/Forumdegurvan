package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"

	_ "github.com/mattn/go-sqlite3"
)

func DeleteLike(l *structures.LikeSentByTheFront) error {
	//establish connection with the database
	db, err := sql.Open("sqlite3", "database/db.sqlite")
	if err != nil {
		return err
	}
	//execute the query
	_, err = db.Exec("DELETE FROM Like WHERE id_post=? AND id_user=?", l.IdPost, l.IdUser)
	if err != nil {
		return err
	}
	return nil
}
