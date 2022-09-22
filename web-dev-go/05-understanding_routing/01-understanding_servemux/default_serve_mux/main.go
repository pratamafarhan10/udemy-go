package main

import (
	"fmt"
	"net/http"
)

type cat struct{}
type dog struct{}

func (c cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>MEOW MEOW MEOW</h1>")
}

func (d dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1>WOOF WOOF WOOF</h1>")
}

func main() {
	var c cat
	var d dog

	http.Handle("/cat", c)
	http.Handle("/dog/", d)

	http.ListenAndServe(":8080", nil)
}
