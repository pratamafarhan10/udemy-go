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
	http.HandleFunc("/", Index)
	http.HandleFunc("/dog/", Dog)
	http.HandleFunc("/me/", Me)

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
