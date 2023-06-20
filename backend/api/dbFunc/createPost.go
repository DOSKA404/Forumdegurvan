package dbFunc

import (
	"database/sql"
	"errors"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	_ "github.com/mattn/go-sqlite3"
)

func CreatePost(post *structures.Post) error {
	db, err := sql.Open("sqlite3", "database/db.sqlite")
	if err != nil {
		e := errors.New("error creating post")
		return e
	}
	q, err := db.Prepare("INSERT INTO Post(date_post, content, id_user) VALUES(?, ?, ?)")
	if err != nil {
		e := errors.New("error creating post")
		return e
	}
	_, err = q.Exec(post.Date_post, post.Content, post.Id_user)
	if err != nil {
		e := errors.New("error creating post")
		return e
	}
	return nil
}
