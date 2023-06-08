package dbFunc

import (
	"database/sql"

	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func GetLike(id_post int) int {
	result := 0
	db, err := sql.Open("sqlite3", "database")
	tools.HandlerError(err)
	row := db.QueryRow("SELECT count(id_like) FROM Like WHERE id_post=?", id_post)
	row.Scan(&result)
	return result
}
