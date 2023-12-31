package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/handlers"
)

func StartApi() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Forum backend api"))
	})
	http.HandleFunc("/createUser", handlers.UserCreationHandler)
	http.HandleFunc("/login", handlers.UserLoginHandlers)
	http.HandleFunc("/createPost", handlers.CreatePost)
	http.HandleFunc("/getPosts", handlers.GetPosts)
	http.HandleFunc("/getUserInfo", handlers.FetchUserInfo)
	http.HandleFunc("/like", handlers.LikePost)
	http.HandleFunc("/unlike", handlers.UnlikePost)
	http.HandleFunc("/changeUsername", handlers.ChangeUsername)
	fmt.Println("api endpoint listening on: http://0.0.0.0:6969")
	log.Fatal(http.ListenAndServe(":6969", nil))
}
