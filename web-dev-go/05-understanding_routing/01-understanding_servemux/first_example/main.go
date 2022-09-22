package main

import (
	"fmt"
	"net/http"
)

type clientHandler struct{}

func (c clientHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/dog":
		fmt.Fprintln(w, "<h1>WOOF WOOF WOOF</h1>")
	case "/cat":
		fmt.Fprintln(w, "<h1>MEOW MEOW MEOW</h1>")
	default:
		fmt.Fprintln(w, "<h1>Page not found</h1>")

	}
}

func main() {
	var ch clientHandler
	http.ListenAndServe(":8080", ch)
}
