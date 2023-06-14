package server

import (
	"fmt"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/handlers"
)

func StartServer() {

	styles := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets", styles)) //We link the css with http.Handle

	http.HandleFunc("/", handlers.MainPage)
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/registerResponse", handlers.RegisterResponse)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/LoginResponse", handlers.LoginResponse)
	http.HandleFunc("/info", handlers.Info)
	http.HandleFunc("/createPost", handlers.CreatePosthandler)
	http.HandleFunc("/test", handlers.Test)
	http.HandleFunc("/testpost", handlers.Testpost)
	http.HandleFunc("/disconnect", handlers.Disconnect)
	http.HandleFunc("/account", handlers.Disconnect)
	//We choose port 80
	fmt.Println("The serveur start on port 8080 🔥") //We print this when the server is online
	fmt.Println("http://localhost:8080/")
	fmt.Println(http.ListenAndServe(":8080", nil)) //We start the server on 80
}
