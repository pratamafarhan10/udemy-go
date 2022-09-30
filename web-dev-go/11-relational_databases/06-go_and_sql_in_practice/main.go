package main

import (
	"net/http"

	"github.com/udemy-go/web-dev-go/11-relational_databases/06-go_and_sql_in_practice/database"
	"github.com/udemy-go/web-dev-go/11-relational_databases/06-go_and_sql_in_practice/handler"
)

func main() {
	defer database.Db.Close()
	http.HandleFunc("/", handler.Index)
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	http.HandleFunc("/logout", handler.Logout)
	http.HandleFunc("/user", handler.DefaultUser)
	http.HandleFunc("/admin", handler.Admin)
	http.ListenAndServe(":8080", nil)
}
