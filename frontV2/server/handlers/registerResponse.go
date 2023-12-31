package handlers

import (
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/apiCall"
	"github.com/GurvanN22/Forum/src/front/server/data"
	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func RegisterResponse(w http.ResponseWriter, r *http.Request) {
	structure := structures.User{
		Email:       r.FormValue("email"),
		Password:    r.FormValue("psw"),
		DateOfBirth: r.FormValue("date"),
		Username:    r.FormValue("pseudo"),
	}
	data.RegisterResponse = apiCall.CreateAccount(&structure)
	switch data.RegisterResponse {
	case "user already exists":
		Register(w, r)
		break
	case "username already exists":
		Register(w, r)
		break
	default:
		Login(w, r)
	}
}
