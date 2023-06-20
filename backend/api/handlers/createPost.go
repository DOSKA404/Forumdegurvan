package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/dbFunc"
	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		byteReq, err := io.ReadAll(r.Body)
		tools.HandlerError(w, err)
		var post structures.Post
		err = json.Unmarshal(byteReq, &post)
		tools.HandlerError(w, err)
		err = dbFunc.CreatePost(&post)
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write([]byte("Post Created"))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request Method"))
	}
}
