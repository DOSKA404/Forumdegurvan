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
		Username:    r.FormValue("pseudo"),
		DateOfBirth: r.FormValue("date"),
	}
	data.RegisterResponse = apiCall.CreateAccount(&structure)
	switch data.RegisterResponse {
	case "user already exists":
		Register(w, r)
		break
	case "username already exists":
		Register(w, r)
		break
	}
}
