package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/read", read)
	http.HandleFunc("/setmore", setMore)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "first-cookie",
		Value: "this is the first cookie",
	})
	io.WriteString(w, "first cookie is written")
}

func read(w http.ResponseWriter, req *http.Request) {
	fc, err := req.Cookie("first-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, fc.String())
	}

	sc, err := req.Cookie("second-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, sc.String())
	}

	tc, err := req.Cookie("third-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, tc.String())
	}
}

func setMore(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "second-cookie",
		Value: "this is the second cookie",
	})
	http.SetCookie(w, &http.Cookie{
		Name:  "third-cookie",
		Value: "this is the third cookie",
	})
	io.WriteString(w, "second and third cookie is written")
}
