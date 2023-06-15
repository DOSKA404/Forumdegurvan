package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/dbFunc"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		listPost, err := dbFunc.GetPosts()
		fmt.Println(listPost)
		tools.HandlerError(w, err)
		byteListPost, err := json.Marshal(&listPost)
		tools.HandlerError(w, err)
		w.Write(byteListPost)
	} else {
		w.Write([]byte("Bad request method"))
	}
}
