package handlers

import (
	"html/template"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/apiCall"
	"github.com/GurvanN22/Forum/src/front/server/data"
	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func Info(w http.ResponseWriter, r *http.Request) {
	User := apiCall.GetUserInfos(data.Email_User)
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
