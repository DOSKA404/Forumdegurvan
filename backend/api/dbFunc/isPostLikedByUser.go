package dbFunc

import "database/sql"

func IsPostLikedByUser(idUser, idPost int, db *sql.DB) bool {
	var result int = 0
	row := db.QueryRow("SELECT COUNT(id_like) FROM Like WHERE id_user=? AND id_post=?", idUser, idPost)
	row.Scan(result)
	return result == 1
}
