package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type response struct {
	Auth    bool
	Message string
	Body    map[string][]string
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/user", user)
	router.GET("/dashboard", dashboard)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res := response{
		Auth:    false,
		Message: "This is an index",
		Body:    nil,
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Fprintln(w, string(b))
}

func isAuthenticated(username, password string) bool {
	if username == "windahbasudara" && password == "password" {
		return true
	}
	return false
}

func auth(w http.ResponseWriter, username, password string) {
	ia := isAuthenticated(username, password)
	if !ia {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
	}
}

func user(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	auth(w, "windahbasudara", "password")
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err.Error())
	}

	res := response{
		Auth:    true,
		Message: "This is a user",
		Body:    req.Form,
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Fprintln(w, string(b))
}

func dashboard(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	auth(w, "ww", "password")
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err.Error())
	}

	res := response{
		Auth:    true,
		Message: "This is a dashboard",
		Body:    req.PostForm,
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Fprintln(w, string(b))
}
