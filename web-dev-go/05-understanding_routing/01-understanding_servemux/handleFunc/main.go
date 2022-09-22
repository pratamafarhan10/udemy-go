package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/cat", catHandler)
	http.HandleFunc("/dog/", dogHandler)

	fmt.Printf("%T\n", catHandler)

	http.ListenAndServe(":8080", nil)
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>WOOF WOOF WOOF</h1>")
}

func catHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "<h1>MEOW MEOW MEOW</h1>")
}
