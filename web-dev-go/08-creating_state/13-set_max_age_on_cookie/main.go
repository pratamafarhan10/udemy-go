package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<h1><a href="/set">Set Cookie</a></h1>
	`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "my-cookie",
		Value:  "this is a cookie",
		MaxAge: 30,
	})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<h1><a href="/read">Read your cookie</a></h1>
	`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<h1><a href="/delete">Delete your cookie</a></h1>
	<br>
	<h1>`+c.String()+`</h1>`)
}

func delete(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-cookie")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1

	http.SetCookie(w, c)
	io.WriteString(w, `
	<h1><a href="/">Back to index</a></h1>
	`)
}
