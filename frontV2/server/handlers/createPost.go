package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/GurvanN22/Forum/src/front/server/apiCall"
	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func CreatePosthandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreatePosthandler")
	IdUser, err := strconv.Atoi(r.FormValue("authorid"))
	if err != nil {
		fmt.Println(err)
	}
	Data := structures.Post{
		Id_user:      IdUser,
		Content:      r.FormValue("content"),
		CreationDate: r.FormValue("date"),
	}
	apiCall.CreatePost(&Data)
	fmt.Println(Data)
	//Info(w, r)
}
