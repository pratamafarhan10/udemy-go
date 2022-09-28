package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	var s string

	if req.Method == http.MethodPost {
		f, h, err := req.FormFile("q")
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		b, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		s = string(b)

		fmt.Println(h.Filename, h.Header, h.Size)

		nf, err := os.Create(h.Filename)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		_, err = nf.Write(b)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="post" enctype="multipart/form-data">
		<input type="file" name="q">
		<input type="submit">
	</form>
	`+s)
}
