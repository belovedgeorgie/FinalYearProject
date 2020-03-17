package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ToDo: initially parse html template one time only

func LoginPage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// use (%%) instead of {{}} for templates
	tmpl, terr := template.New("login.html").Delims("(%", "%)").ParseFiles("views/loginpage/login.html")
	if terr != nil {
		log.Fatalln(terr)
	}

	if err := tmpl.ExecuteTemplate(w, "login.html", nil); err != nil {
		log.Println(err)
	}
}
