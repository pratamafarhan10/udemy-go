package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/dog/", Dog)
	http.HandleFunc("/me/", Me)

	http.ListenAndServe(":8080", nil)
}

func Index(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Welcome")
}

func Dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "WOOF WOOF WOOF")
}

func Me(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Windah Basudara")
}
