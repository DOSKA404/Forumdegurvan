package apiCall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/data"
	"github.com/GurvanN22/Forum/src/front/server/structures"
	//"github.com/GurvanN22/Forum/src/front/server/tools"
)

func CreateAccount(structure *structures.User) string {

	//hash := tools.HashToken(structure.Email + structure.Password)
	bytesStructure, err := json.Marshal(structure)
	if err != nil {
		fmt.Println(err)
	}
	req, err := http.NewRequest(http.MethodPost, data.APIadress+"/createUser", bytes.NewReader(bytesStructure))
	if err != nil {
		fmt.Println(err)
	}
	//req.Header.Add("Token :", string(hash))
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
