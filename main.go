package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/metaclips/FinalYearProject/controller"
)

func main() {
	router := httprouter.New()

	router.GET("/", controller.HomePage)
	router.GET("/login", controller.LoginPage)

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalln(err)
	}
}
