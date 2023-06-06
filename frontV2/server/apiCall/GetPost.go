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

func GetPost() structures.Post {
	//hash := tools.HashToken(structure.Email + structure.Password)

	req, err := http.NewRequest(http.MethodGet, "http://localhost:6969/getPosts", bytes.NewReader([]byte("")))
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
	Posts := structures.Post{}
	json.Unmarshal(byteRes, &Posts)
	return Posts
}
