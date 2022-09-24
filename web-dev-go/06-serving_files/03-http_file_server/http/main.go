package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	io.WriteString(w, `<h1>Hello World</h1>`)
// }

func dog(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./files/dog.jpg")
}
