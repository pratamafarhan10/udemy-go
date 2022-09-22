package main

import (
	"fmt"
	"net/http"
)

type cat struct{}
type dog struct{}

func (c cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>MEOW MEOW MEOW</h1>")
}

func (d dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>WOOF WOOF WOOF</h1>")
}

func main() {
	var c cat
	var d dog

	mux := http.NewServeMux()
	mux.Handle("/cat", c)  // only serve /cat
	mux.Handle("/dog/", d) // can serve /dog/this/too

	http.ListenAndServe(":8080", mux)
}
