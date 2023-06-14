package handlers

import (
	"html/template"
	"net/http"
)

func Disconnect(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w, r)
	tmpl := template.Must(template.ParseFiles("html/loginre.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}

func DeleteCookie(w http.ResponseWriter, r *http.Request) {
	Username := http.Cookie{
		Name:     "Username",
		Value:    "",
		Path:     "/",
		MaxAge:   1704085200,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	Email := http.Cookie{
		Name:     "EmailUser",
		Value:    "",
		Path:     "/",
		MaxAge:   1704085200,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	ID := http.Cookie{
		Name:     "IdUser",
		Value:    "",
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
