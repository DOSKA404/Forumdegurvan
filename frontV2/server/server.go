package server

import (
	"fmt"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/handlers"
)

func StartServer() {

	styles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", styles)) //We link the css with http.Handle

	http.HandleFunc("/", handlers.Login)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/registerResponse", handlers.RegisterResponse)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/LoginResponse", handlers.LoginResponse)
	http.HandleFunc("/info", handlers.Info)
	http.HandleFunc("/createPost", handlers.CreatePosthandler)
	//We choose port 8080
	fmt.Println("The serveur start on port 8080 🔥") //We print this when the server is online
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil) //We start the server on 8080

}
<<<<<<< HEAD
func page(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/main.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
func register(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/createAccount.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
=======
>>>>>>> main
