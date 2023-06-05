package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func CreatePost(post *structures.Post) {
	db, err := sql.Open("sqlite3", "database/db.db")
	tools.HandlerError(err)
	q, err := db.Prepare("INSERT INTO Post(date_post, content, id_user) VALUES(?, ?, ?)")
	tools.HandlerError(err)
	_, err = q.Exec(post.Date_post, post.Content, post.Id_user)
	tools.HandlerError(err)
}
