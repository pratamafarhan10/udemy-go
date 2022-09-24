package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/pdf", pdf)

	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/dog.jpg">`)
}

func dog(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("dog.jpg")
	if err != nil {
		http.Error(res, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	io.Copy(res, f)
}

func pdf(res http.ResponseWriter, req *http.Request) {
	f, err := os.Open("pdf.pdf")
	if err != nil {
		http.Error(res, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	io.Copy(res, f)
}
