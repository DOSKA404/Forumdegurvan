package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/GurvanN22/Forum/src/Backend/api/dbFunc"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		byteBody, err := io.ReadAll(r.Body)
		tools.HandlerError(w, err)
		tmp := string(byteBody)
		idUser, err := strconv.Atoi(tmp)
		tools.HandlerError(w, err)
		listPost, err := dbFunc.GetPosts(idUser)
		tools.HandlerError(w, err)
		byteListPost, err := json.Marshal(&listPost)
		tools.HandlerError(w, err)
		w.Write(byteListPost)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request Method"))
	}
}
