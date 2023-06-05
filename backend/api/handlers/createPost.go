package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		byteReq, err := io.ReadAll(r.Body)
		tools.HandlerError(err)
		var post structures.Post
		err = json.Unmarshal(byteReq, &post)
		tools.HandlerError(err)
		
	} else {
		w.Write([]byte("bad request method"))
	}
}
