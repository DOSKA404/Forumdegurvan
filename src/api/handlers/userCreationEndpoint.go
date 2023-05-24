package handlers

import "net/http"

func UserCreationHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("user creation endpoint"))
}