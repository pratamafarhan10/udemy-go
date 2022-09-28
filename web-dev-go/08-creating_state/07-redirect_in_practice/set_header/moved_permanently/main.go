package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/success", success)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL, req.Method)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>Index</h1>")
}

func success(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL, req.Method)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Location", "/success")
	w.WriteHeader(http.StatusMovedPermanently)
	io.WriteString(w, "<h1>Success</h1>")
}
