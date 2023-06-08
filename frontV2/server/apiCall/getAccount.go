package apiCall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/structures"
	//"github.com/GurvanN22/Forum/src/front/server/tools"
)

func GetUserInfos(email string) structures.User {
	//hash := tools.HashToken(structure.Email + structure.Password)
	req, err := http.NewRequest(http.MethodPost, "http://localhost:6969/getUserInfo", bytes.NewReader([]byte(email)))
	if err != nil {
		fmt.Println(err)
	}
	//req.Header.Add("Token", string(hash))
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	byteRes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	User := structures.User{}
	json.Unmarshal(byteRes, &User)
	return User
}
