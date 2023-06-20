package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/dbFunc"
	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func LikePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		byteBody, err := io.ReadAll(r.Body)
		tools.HandlerError(w, err)
		var likeSendByTheFront structures.LikeSentByTheFront
		err = json.Unmarshal(byteBody, &likeSendByTheFront)
		tools.HandlerError(w, err)
		dbFunc.CreateLike(&likeSendByTheFront)
		w.Write([]byte("Like created"))
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request Method"))
	}
}
