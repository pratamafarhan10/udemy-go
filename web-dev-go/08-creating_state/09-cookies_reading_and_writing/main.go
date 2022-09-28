package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/foo", foo)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "this is a cookie",
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>Coba Cookie</h1>")
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, "cookie not found", http.StatusNotFound)
		return
	}
	fmt.Println("Cookie string:", c.String())
	fmt.Println("Name", c.Name, "Value", c.Value)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<h1>foo</h1>")
}
