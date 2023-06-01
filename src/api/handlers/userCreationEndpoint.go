package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/dbFunc"
	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func UserCreationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodPost {
		var user *structures.User
		byteUser, err := io.ReadAll(r.Body)
		tools.HandlerError(err)
		err = json.Unmarshal(byteUser, &user)
		tools.HandlerError(err)
		err = dbFunc.PutUserInDb(user)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("user created"))
		}
	} else {
		w.Write([]byte("Bad request Method"))
	}
}
