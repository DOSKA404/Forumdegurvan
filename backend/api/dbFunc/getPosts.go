package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func GetPosts() []structures.PostWithLike {
	db, err := sql.Open("sqlite3", "database/db.db")
	tools.HandlerError(err)
	defer db.Close()
	rows, err := db.Query("SELECT Post.id_post, Post.content, User.Username, User.id_user FROM Post INNER JOIN User ON Post.id_user=User.id_user")
	tools.HandlerError(err)
	result := []structures.PostWithLike{}
	for rows.Next() {
		var post structures.PostWithLike
		err = rows.Scan(&post.Id_post, &post.Content, &post.Username, &post.Id_user)
		tools.HandlerError(err)
		result = append(result, post)
	}
	return result
}
