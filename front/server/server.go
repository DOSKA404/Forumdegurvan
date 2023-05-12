package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	styles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", styles)) //We link the css with http.Handle
	http.HandleFunc("/", page)                                   //We create the main page , the only function who use a template

	Port := "8080"                                          //We choose port 8080
	fmt.Println("The serveur start on port " + Port + " 🔥") //We print this when the server is online
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":"+Port, nil) //We start the server}

}
func page(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./index.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
