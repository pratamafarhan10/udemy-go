package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/form", form)
	http.HandleFunc("/success", success)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL, req.Method)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>Index</h1>")
}

func form(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		fmt.Println(req.URL, req.Method)
		w.Header().Set("Location", "/success")
		w.WriteHeader(http.StatusSeeOther)
		return
	}
	fmt.Println(req.URL, req.Method)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="POST" action="/form">
		<input type="text" name="q">
		<input type="submit">
	</form>
	`)
}

func success(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL, req.Method)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>Success</h1>")
}
