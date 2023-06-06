package handlers

import (
	"fmt"
	"net/http"
)

func Testpost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}
