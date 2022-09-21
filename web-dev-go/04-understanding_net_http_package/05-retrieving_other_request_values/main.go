package main

import (
	"log"
	"net/http"
	"net/url"
	"text/template"
)

var tpl *template.Template

type clientHandler struct{}

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var ch clientHandler
	http.ListenAndServe(":8080", ch)
}

func (c clientHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err.Error())
	}

	data := struct {
		Method                  string
		URL                     *url.URL
		Submissions             map[string][]string
		Header                  map[string][]string
		Host                    string
		ContentLength           int64
		Name, Phone_number, Age string
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
		req.FormValue("name"),
		req.FormValue("phone_number"),
		req.FormValue("age"),
	}

	err = tpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
