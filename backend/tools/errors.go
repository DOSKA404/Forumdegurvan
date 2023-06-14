package tools

import (
	"net/http"
)

func HandlerError(w http.ResponseWriter, e error) {
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error api"))
	}
}
