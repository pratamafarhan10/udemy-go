package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	// router.GET("/", index)
	router.Handler("GET", "/", http.FileServer(http.Dir(".")))
	router.GET("/file/:filename", dog)
	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<h1>Hello World</h1>`)
}

func dog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	http.ServeFile(w, r, "./files/"+p.ByName("filename"))
}
