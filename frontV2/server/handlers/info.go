package handlers

import (
	"html/template"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func Info(w http.ResponseWriter, r *http.Request) {
	Data := structures.InfoDataInput{
		USERNAME: "Le Rouge",
		PostLine: []structures.Post{},
		IdUser:   666,
	}
	i := 5
	for i > 0 {
		Data.PostLine = append(Data.PostLine, structures.Post{
			Username: "Le Rouge",
			Content:  "le front est chiant",
			Likes:    666,
			IsLike:   i%2 == 0,
		})
		i--
	}
	tmpl := template.Must(template.ParseFiles("html/infoFile.html")) //We link the template and the html file
	tmpl.Execute(w, Data)
	tmpl = template.Must(template.ParseFiles("html/footer.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
