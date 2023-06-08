package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/GurvanN22/Forum/src/front/server/apiCall"
	"github.com/GurvanN22/Forum/src/front/server/data"
	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func Info(w http.ResponseWriter, r *http.Request) {
	User := SetCookies(w, r)
	Data := structures.InfoDataInput{
		USERNAME: User.Username,
		PostLine: []structures.Post{},
		IdUser:   User.Id_user,
	}
	Data.PostLine = apiCall.GetPost()
	tmpl := template.Must(template.ParseFiles("html/infoFile.html")) //We link the template and the html file
	tmpl.Execute(w, Data)
	tmpl = template.Must(template.ParseFiles("html/footer.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}

func SetCookies(w http.ResponseWriter, r *http.Request) structures.User {
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
	return User
}
