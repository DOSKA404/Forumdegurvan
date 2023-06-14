package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/GurvanN22/Forum/src/front/server/apiCall"
	"github.com/GurvanN22/Forum/src/front/server/structures"
	"github.com/GurvanN22/Forum/src/front/server/tools"
)

func Info(w http.ResponseWriter, r *http.Request) {
	Username, err := r.Cookie("Username")
	tools.CheckErr(err)
	IdUser, err := r.Cookie("IdUser")
	tools.CheckErr(err)
	IdUserint, err := strconv.Atoi(IdUser.Value)
	tools.CheckErr(err)
	Data := structures.InfoDataInput{
		USERNAME: Username.Value,
		PostLine: []structures.Post{},
		IdUser:   IdUserint,
	}
	Data.PostLine = apiCall.GetPost()
	tmpl := template.Must(template.ParseFiles("html/header.html")) //We link the template and the html file
	tmpl.Execute(w, Data)
	tmpl = template.Must(template.ParseFiles("html/infoFile.html")) //We link the template and the html file
	tmpl.Execute(w, Data)
	tmpl = template.Must(template.ParseFiles("html/footer.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
