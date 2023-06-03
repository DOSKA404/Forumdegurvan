package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/GurvanN22/Forum/src/front/server/data"
	"github.com/GurvanN22/Forum/src/front/server/structures"
)

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println(data.RegisterResponse)
	switch data.RegisterResponse {
	case "user already exists":
		Data := structures.Error{
			ErrorMSG: data.RegisterResponse,
			IsError:  true,
			IsError2: false,
		}
		tmpl := template.Must(template.ParseFiles("html/createAccount.html")) //We link the template and the html file
		tmpl.Execute(w, Data)
		break
	case "username already exists":
		fmt.Println("username already exists inside")
		Data := structures.Error{
			ErrorMSG: data.RegisterResponse,
			IsError:  false,
			IsError2: true,
		}
		tmpl := template.Must(template.ParseFiles("html/createAccount.html")) //We link the template and the html file
		tmpl.Execute(w, Data)
		break
	default:
		tmpl := template.Must(template.ParseFiles("html/createAccount.html")) //We link the template and the html file
		tmpl.Execute(w, nil)
		break

	}
	tmpl := template.Must(template.ParseFiles("html/footer.html")) //We link the template and the html file
	tmpl.Execute(w, nil)
}
