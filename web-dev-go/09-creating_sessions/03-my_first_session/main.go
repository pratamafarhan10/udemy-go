package main

import (
	"io"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	username, password, firstname, lastname string
	age                                     int
}

var dbUser = map[string]user{ // User ID, User
	"a90sdf": {
		username:  "windah",
		password:  "password",
		firstname: "windah",
		lastname:  "basudara",
		age:       30,
	},
	"a0asko": {
		username:  "james",
		password:  "password",
		firstname: "james",
		lastname:  "bond",
		age:       42,
	},
}

var dbSession = map[string]string{} // session ID, user ID

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err != nil {
		io.WriteString(w, `
		<h1>You are not authenticated</h1>
		<a href="/login">Login</a>
		`)
		return
	}
	userId, ok := dbSession[c.Value]
	if !ok {
		http.Redirect(w, req, "/logout", http.StatusSeeOther)
		return
	}
	fn := dbUser[userId].firstname
	io.WriteString(w, `
	<h1>Welcome to Dashboard</h1>
	<h2>Good morning `+fn+`<br><a href="/logout">Logout</a>
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
			dbSession[s.String()] = "a90sdf"
			http.Redirect(w, req, "/", http.StatusSeeOther)
			return
		} else if u == "james" && p == "password" {
			s := uuid.New()
			cookie := &http.Cookie{
				Name:     "session",
				Value:    s.String(),
				HttpOnly: true,
				MaxAge:   60 * 10,
			}
			http.SetCookie(w, cookie)
			dbSession[s.String()] = "a0asko"
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
