package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/03-mvc_design_pattern/controllers"
)

func main() {
	router := httprouter.New()
	user := controllers.NewUserController()
	router.GET("/user/:id", user.GetUser)
	router.POST("/user", user.CreateUser)
	log.Fatalln(http.ListenAndServe(":8080", router))
}
