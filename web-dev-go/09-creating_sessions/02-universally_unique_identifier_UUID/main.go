package main

import (
	"io"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("session")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		io.WriteString(w, `
		<h1>You are not authenticated</h1>
		<a href="/login">Login</a>
		`)
		return
	}
	io.WriteString(w, `
	<h1>Welcome to Dashboard</h1>
	<a href="/logout">Logout</a>
	`)
}

func login(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("session")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err == nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var errMessage string

	if req.Method == http.MethodPost {
		u := req.FormValue("username")
		p := req.FormValue("password")

		if u == "windah" && p == "password" {
			s := uuid.New()
			cookie := &http.Cookie{
				Name:     "session",
				Value:    s.String(),
				HttpOnly: true,
				MaxAge:   60 * 10,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		} else {
			errMessage = "Username or password wrong"
		}
	}

	io.WriteString(w, `
	<h1>Login</h1><br>`+errMessage+`<form method="POST">
	<input type="text" name="username">
	<input type="password" name="password">
	<button type="submit">Login</button>
	</form>
	`)
}

func logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	c.MaxAge = -1

	http.SetCookie(w, c)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
