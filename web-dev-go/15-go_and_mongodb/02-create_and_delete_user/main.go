package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/01-organizing_code_into_packages/models"
)

func main() {
	router := httprouter.New()
	router.GET("/user/:id", getUser)
	router.POST("/user", createUser)
	router.DELETE("/user", deleteUser)
	http.ListenAndServe(":8080", router)
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	p1 := models.User{
		FirstName: "James",
		LastName:  "Bond",
		Age:       53,
		Id:        p.ByName("id"),
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Panicln(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(bs))
}

func createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var u models.User

	json.NewDecoder(r.Body).Decode(&u)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "User created")
}

func deleteUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User deleted")
}
