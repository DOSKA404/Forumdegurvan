package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/dbFunc"
	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func UserLoginHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user *structures.User
		byteReq, err := io.ReadAll(r.Body)
		tools.HandlerError(err)
		fmt.Println(string(byteReq))
		err = json.Unmarshal(byteReq, user)
		tools.HandlerError(err)
		hashPassInDb, err := dbFunc.GetPasswordWithEmail(user.Email)
		tools.HandlerError(err)
		tools.CheckPasswordValidity(hashPassInDb, user.Password)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	} else {
		w.Write([]byte("Bad request method"))
	}
}
