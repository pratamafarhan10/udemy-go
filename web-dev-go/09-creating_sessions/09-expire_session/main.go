package main

import (
	"net/http"

	"github.com/udemy-go/web-dev-go/09-creating_sessions/09-expire_session/handler"
)

func main() {
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/user", handler.DefaultUser)
	http.HandleFunc("/admin", handler.Admin)
	http.ListenAndServe(":8080", nil)
}
