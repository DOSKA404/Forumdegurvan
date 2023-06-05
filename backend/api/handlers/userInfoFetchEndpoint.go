package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/dbFunc"
	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func FetchUserInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var username string
		var user structures.User

		body, err := io.ReadAll(r.Body)
		tools.HandlerError(err)

		err = json.Unmarshal(body, &user)
		tools.HandlerError(err)

		username, err = dbFunc.GetUsername(user.Email)
		user.Username = username
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			byteUser, err := json.Marshal(&user)
			tools.HandlerError(err)
			w.Write(byteUser)
		}
	} else {
		w.Write([]byte("Bad request Method"))
	}
}
