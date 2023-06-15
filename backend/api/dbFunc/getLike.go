package dbFunc

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func GetLike(id_post int) (int, error) {
	result := 0
	db, err := sql.Open("sqlite3", "database")
	if err != nil {
		return 0, err
	}
	row := db.QueryRow("SELECT count(id_like) FROM Like WHERE id_post=?", id_post)
	fmt.Println(id_post)
	row.Scan(&result)
	fmt.Println(result)
	return result, nil
}
