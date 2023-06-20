package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	_ "github.com/mattn/go-sqlite3"
)

func CreateLike(likeSendByTheFront *structures.LikeSentByTheFront) error {
	db, err := sql.Open("sqlite3", "database/db.sqlite")
	if err != nil {
		return err
	}
	req, err := db.Prepare("INSERT INTO Like(id_post, id_user) VALUES(?, ?)")
	if err != nil {
		return err
	}
	_, err = req.Exec(&likeSendByTheFront.IdPost, &likeSendByTheFront.IdUser)
	if err != nil {
		return err
	}
	return nil
}
