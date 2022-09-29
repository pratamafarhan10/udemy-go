package main

import (
	"net/http"

	"github.com/udemy-go/web-dev-go/09-creating_sessions/04-sign_up/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/logout", handler.Logout)
	http.ListenAndServe(":8080", nil)
}
