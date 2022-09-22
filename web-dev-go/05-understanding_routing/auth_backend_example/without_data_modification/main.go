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
	router.GET("/user", auth(user, "windahbasudara", "password"))
	router.GET("/dashboard", auth(dashboard, "ww", "password"))

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

func auth(h httprouter.Handle, username, password string) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		if username == "windahbasudara" && password == "password" {
			h(w, req, p)
		} else {
			http.Error(w, "user not authenticated", http.StatusUnauthorized)
		}
	}
}

func user(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	req.ParseForm()
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
	res := response{
		Auth:    true,
		Message: "This is a dashboard",
		Body:    nil,
	}

	b, err := json.Marshal(res)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Fprintln(w, string(b))
}
