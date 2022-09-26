package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog", dog)
	http.Handle("/resource/", http.StripPrefix("/resource", http.FileServer(http.Dir("./assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resource/dog.jpg">`)
}

func dog(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("./assets/dog.jpg")
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
	}

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", http.StatusNotFound)
	}

	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
