package handlers

import (
	"html/template"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/structures"
	"github.com/GurvanN22/Forum/src/front/server/apiCall"
)

func Info(w http.ResponseWriter, r *http.Request) {
	Data := structures.InfoDataInput{
		USERNAME: "Le Rouge",
		PostLine: []structures.Post{},
		IdUser:   1,
	}
	Data.PostLine = apiCall.GetPost()
	tmpl := template.Must(template.ParseFiles("html/infoFile.html")) //We link the template and the html file
	tmpl.Execute(w, Data)
	tmpl = template.Must(template.ParseFiles("html/footer.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
