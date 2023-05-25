package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/tools"
)

type User struct {
	Email       string
	Username    string
	DateOfBirth string
	Hash        string
}

func UserCreationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var user *User
		byteUser, err := io.ReadAll(r.Body)
		tools.HandlerError(err)
		err = json.Unmarshal(byteUser, &user)
		tools.HandlerError(err)
	} else {
		w.Write([]byte("Bad request Method"))
	}
}
