package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func UserLoginHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user *structures.User
		byteReq, err := io.ReadAll(r.Body)
		tools.HandlerError(err)
		err = json.Unmarshal(byteReq, user)
		tools.HandlerError(err)
	} else {
		w.Write([]byte("Bad request method"))
	}
}
