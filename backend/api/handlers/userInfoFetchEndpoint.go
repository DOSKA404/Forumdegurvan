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
		var userInfo structures.User
		byteEmail, err := io.ReadAll(r.Body)
		tools.HandlerError(err)
		userInfo = dbFunc.GetUserInfo(string(byteEmail))
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			byteUser, err := json.Marshal(&userInfo)
			tools.HandlerError(err)
			w.Write(byteUser)
		}
	} else {
		w.Write([]byte("Bad request Method"))
	}
}