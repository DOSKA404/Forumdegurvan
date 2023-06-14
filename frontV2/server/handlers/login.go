package handlers

import (
	"html/template"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/data"
	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func Login(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("EmailUser")
	if err != nil || cookie.Value == "" {
		if data.LoginResponse == "incorrect password" || data.LoginResponse == "no user found" {
			Data := structures.Error{
				ErrorMSG: data.LoginResponse,
				IsError:  true,
			}
			tmpl := template.Must(template.ParseFiles("html/login.html")) //We link the template and the html file
			tmpl.Execute(w, Data)
			tmpl = template.Must(template.ParseFiles("html/footer.html")) //We link the template and the html file
			tmpl.Execute(w, nil)
		} else {
			tmpl := template.Must(template.ParseFiles("html/login.html")) //We link the template and the html file
			tmpl.Execute(w, nil)
			tmpl = template.Must(template.ParseFiles("html/footer.html")) //We link the template and the html file
			tmpl.Execute(w, nil)
		}
		data.LoginResponse = ""
	} else {
		data.Email_User = cookie.Value
		Info(w, r)
	}

}
