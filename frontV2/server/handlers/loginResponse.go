package handlers

import (
	"fmt"
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
	switch data.RegisterResponse {
	case "Login successful":
		fmt.Println("Login successful")
		Login(w, r)
		break
	case "incorrect password":
		fmt.Println("incorrect password")
		Login(w, r)
		break
	case "no user found":
		fmt.Println("no user found")
		break
	}
}
