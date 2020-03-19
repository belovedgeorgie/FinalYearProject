package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ToDo: initially parse html template one time only

func HomePage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// use (%%) instead of {{}} for templates
	tmpl, terr := template.New("home.html").Delims("(%", "%)").ParseFiles("views/homepage/home.html", "views/homepage/components/SideBar.vue")
	if terr != nil {
		log.Fatalln(terr)
	}
	if err := tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		log.Println(err)
	}
}
