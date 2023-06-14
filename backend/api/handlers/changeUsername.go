package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/Backend/api/dbFunc"
	"github.com/GurvanN22/Forum/src/Backend/api/structures"
	"github.com/GurvanN22/Forum/src/Backend/tools"
)

func ChangeUsername(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		var u structures.UpdateUsername
		byteBody, err := io.ReadAll(r.Body)
		if err != nil {
			tools.HandlerError(w, err)
		}
		err = json.Unmarshal(byteBody, &u)
		if err != nil {
			tools.HandlerError(w, err)
		}
		err = dbFunc.UpdateUsername(&u)
		if err != nil {
			tools.HandlerError(w, err)
		} else {
			w.Write([]byte("Username successfully updated"))
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request Method"))
	}
}
