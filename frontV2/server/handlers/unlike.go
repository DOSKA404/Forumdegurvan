package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/GurvanN22/Forum/src/front/server/apiCall"
	"github.com/GurvanN22/Forum/src/front/server/structures"
	"github.com/GurvanN22/Forum/src/front/server/tools"
)

func Unlike(w http.ResponseWriter, r *http.Request) {
	idpost, err := strconv.Atoi(r.FormValue("id-post"))
	tools.CheckErr(err)
	IdUser, err := r.Cookie("IdUser")
	tools.CheckErr(err)
	IdUserint, err := strconv.Atoi(IdUser.Value)
	tools.CheckErr(err)
	struc := structures.LikeSentByTheFront{
		IdPost: idpost,
		IdUser: IdUserint,
	}
	apiCall.Unlike(&struc)

	tmpl := template.Must(template.ParseFiles("html/infore.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
