package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func CreateLike(likeSendByTheFront *structures.LikeSentByTheFront) {
	db, err := sql.Open("sqlite3", "database/db.db")
	tools.HandlerError(err)
	req, err := db.Prepare("INSERT INTO Like(id_post, id_user) VALUES(?, ?)")
	tools.HandlerError(err)
	_, err = req.Exec(&likeSendByTheFront.IdPost, &likeSendByTheFront.IdUser)
	tools.HandlerError(err)
}
