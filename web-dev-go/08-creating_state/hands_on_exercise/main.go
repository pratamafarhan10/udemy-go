package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visit",
			Value: "0",
		}
	}

	i, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	i += 1

	cookie.Value = strconv.Itoa(i)

	http.SetCookie(w, cookie)

	fmt.Println(cookie.Name, i)
	fmt.Fprintf(w, "This is index page")
}

func contact(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visit",
			Value: "0",
		}
	}

	i, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	i += 1

	cookie.Value = strconv.Itoa(i)

	http.SetCookie(w, cookie)

	fmt.Println(cookie.Name, i)
	fmt.Fprintf(w, "This is contact page")
}

func about(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visit",
			Value: "0",
		}
	}

	i, err := strconv.Atoi(cookie.Value)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	i += 1

	cookie.Value = strconv.Itoa(i)

	http.SetCookie(w, cookie)

	fmt.Println(cookie.Name, i)
	fmt.Fprintf(w, "This is about page")
}
