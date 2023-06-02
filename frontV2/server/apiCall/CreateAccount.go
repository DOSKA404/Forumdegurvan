package apiCall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func CreateAccount(structure *structures.User) string {
	bytesStructure, err := json.Marshal(structure)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest(http.MethodPost, "http://localhost:6969/createUser", bytes.NewReader(bytesStructure))
	if err != nil {
		fmt.Println(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	byteRes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(byteRes)
}
