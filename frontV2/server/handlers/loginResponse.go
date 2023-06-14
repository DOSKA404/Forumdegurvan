package handlers

import (
	"html/template"
	"net/http"
	"strconv"

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
		data.Email_User = structure.Email
		SetCookies(w, r)
		tmpl := template.Must(template.ParseFiles("html/infore.html")) //We link the template and the html file
		tmpl.Execute(w, nil)
		break
	case "incorrect password":
		Login(w, r)
		break
	case "no user found":
		Login(w, r)
		break
	}
}

func SetCookies(w http.ResponseWriter, r *http.Request) {
	User := apiCall.GetUserInfos(data.Email_User)
	Username := http.Cookie{
		Name:     "Username",
		Value:    User.Username,
		Path:     "/",
		MaxAge:   1704085200,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	Email := http.Cookie{
		Name:     "EmailUser",
		Value:    User.Email,
		Path:     "/",
		MaxAge:   1704085200,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	ID := http.Cookie{
		Name:     "IdUser",
		Value:    strconv.Itoa(User.Id_user),
		Path:     "/",
		MaxAge:   1704085200,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &Username)
	http.SetCookie(w, &ID)
	http.SetCookie(w, &Email)
}
