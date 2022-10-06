package main

import (
	"log"
	"net/http"

	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/models"
	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/router"
)

func main() {
	cancel := models.ContextCancel
	defer cancel()
	log.Fatalln(http.ListenAndServe(":8080", router.Routers()))
}
