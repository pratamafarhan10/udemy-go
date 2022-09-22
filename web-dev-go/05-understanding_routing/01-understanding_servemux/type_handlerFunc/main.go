package main

import (
	"fmt"
	"net/http"
)

// var hf http.HandlerFunc

func main() {
	// Using HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	// http.HandleFunc("/cat", catHandler)
	// http.HandleFunc("/dog/", dogHandler)

	// Using Handle(pattern string, handler Handler)
	http.Handle("/cat", http.HandlerFunc(catHandler))
	http.Handle("/dog", http.HandlerFunc(dogHandler))

	fmt.Printf("%T\n", catHandler)

	http.ListenAndServe(":8080", nil)
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>WOOF WOOF WOOF</h1>")
}

func catHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>MEOW MEOW MEOW</h1>")
}
