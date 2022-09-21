package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type clientHandler struct{}

func (c clientHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err.Error())
	}

	// Get all the data including the query parameter
	fmt.Println(r.Form)
	// Get data from body only
	fmt.Println(r.PostForm)
	// Get the method
	fmt.Println(r.Method)
	// Get the url of request
	fmt.Println(r.URL)
	// Get the header
	fmt.Println(r.Header)

	// fmt.Println(r.Body)

	// Close indicates whether to close the connection after replying to this request (for servers) or after sending this request and reading its response (for clients).
	fmt.Println(r.Close)
	// Get the host address
	fmt.Println(r.Host)
	// RemoteAddr allows HTTP servers and other software to record the network address that sent the request, usually for logging.
	fmt.Println(r.RemoteAddr)
	// RequestURI is the targeted URI from client to server
	fmt.Println(r.RequestURI)

	tpl.ExecuteTemplate(w, "index.html", r.Form)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var ch clientHandler
	http.ListenAndServe(":8080", ch)
}
