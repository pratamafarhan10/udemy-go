package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/about", About)
	router.GET("/contact", Contact)
	router.GET("/apply", Apply)
	router.POST("/apply", ApplyPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func About(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "about.html", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.html", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func Apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.html", nil)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func ApplyPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.html", "Apply POST")
	if err != nil {
		log.Fatalln(err.Error())
	}
}
