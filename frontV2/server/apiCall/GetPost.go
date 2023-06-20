package apiCall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/GurvanN22/Forum/src/front/server/data"
	"github.com/GurvanN22/Forum/src/front/server/structures"
	//"github.com/GurvanN22/Forum/src/front/server/tools"
)

func GetPost(id int) []structures.Post {
	//hash := tools.HashToken(structure.Email + structure.Password)
	StringId := strconv.Itoa(id)
	req, err := http.NewRequest(http.MethodGet, data.APIadress+"/getPosts", bytes.NewReader([]byte(StringId)))
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
	Posts := []structures.Post{}
	json.Unmarshal(byteRes, &Posts)
	return Posts
}
