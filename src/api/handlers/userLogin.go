package handlers

import "net/http"

func UserLoginHandlers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Helloooo"))
}
