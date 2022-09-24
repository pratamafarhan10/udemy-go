package main

import (
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", index)
	router.GET("/file/:filename", fileHandler)

	http.ListenAndServe(":8080", router)
}

func index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `<img src="/dog.jpg">`)
}

func fileHandler(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	f, err := os.Open(p.ByName("filename"))
	if err != nil {
		http.Error(res, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	fi, err := f.Stat()

	if err != nil {
		http.Error(res, "failed to get file statistic", http.StatusNotFound)
	}
	http.ServeContent(res, req, f.Name(), fi.ModTime(), f)

}
