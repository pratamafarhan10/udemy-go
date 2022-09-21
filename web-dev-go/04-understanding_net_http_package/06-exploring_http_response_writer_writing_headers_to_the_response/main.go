package main

import (
	"fmt"
	"net/http"
)

type clientHandler struct{}

func (c clientHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Tion-Key", "Wagwan my g")

	// This will return only plain text
	// w.Header().Set("Content-Type", "text/plain")

	// This will return a html
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "<h1>Hello fam</h1>")
}

func main() {
	var ch clientHandler
	http.ListenAndServe(":8080", ch)
}
