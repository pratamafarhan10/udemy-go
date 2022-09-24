package main

import (
	"io"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/file/:filename", fileHandler)

	http.ListenAndServe(":8080", router)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/dog.jpg">`)
}

func fileHandler(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	http.ServeFile(w, req, p.ByName("filename"))
}
