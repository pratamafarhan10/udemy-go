package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3307)/coba_golang_1")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "Successfully Connected")
	check(err)
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
