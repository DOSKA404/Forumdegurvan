package handlers

import (
	"html/template"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/test.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
