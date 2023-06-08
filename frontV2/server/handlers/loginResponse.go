package handlers

import (
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/apiCall"
	"github.com/GurvanN22/Forum/src/front/server/data"
	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func LoginResponse(w http.ResponseWriter, r *http.Request) {
	structure := structures.User{
		Email:    r.FormValue("email"),
		Password: r.FormValue("psw"),
	}
	data.LoginResponse = apiCall.LoginConnection(&structure)
	switch data.LoginResponse {
	case "Login successful":
		Info(w, r)
		break
	case "incorrect password":
		Login(w, r)
		break
	case "no user found":
		Login(w, r)
		break
	}
}
