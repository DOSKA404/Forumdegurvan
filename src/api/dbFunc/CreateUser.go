package dbFunc

import (
	"database/sql"
	"fmt"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
	_ "github.com/mattn/go-sqlite3"
)

func PutUserInDb(user *structures.User) {
	db, err := sql.Open("sqlite3", "database/db.db")
	tools.HandlerError(err)
	fmt.Println(db)
}
