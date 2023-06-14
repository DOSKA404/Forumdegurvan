package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/GurvanN22/Forum/src/front/server/apiCall"
	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func CreatePosthandler(w http.ResponseWriter, r *http.Request) {
	IdUser, err := strconv.Atoi(r.FormValue("authorid"))
	if err != nil {
		fmt.Println(err)
	}
	Data := structures.Post{
		Id_user:   IdUser,
		Content:   r.FormValue("content"),
		Date_post: r.FormValue("date"),
	}
	apiCall.CreatePost(&Data)
	tmpl := template.Must(template.ParseFiles("html/infore.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
