package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
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

		fmt.Println(h.Filename, h.Header, h.Size)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			s += scanner.Text()
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
