package main

import (
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	http.Handle("/", http.HandlerFunc(Index))
	http.Handle("/dog/", http.HandlerFunc(Dog))
	http.Handle("/me/", http.HandlerFunc(Me))

	http.ListenAndServe(":8080", nil)
}

func Index(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.html", "Welcome")
}

func Dog(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.html", "WOOF WOOF WOOF")
}

func Me(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.html", "Windah Basudara")
}
