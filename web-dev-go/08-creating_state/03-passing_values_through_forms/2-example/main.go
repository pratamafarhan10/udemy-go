package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fn := req.FormValue("fn")
	ln := req.FormValue("ln")
	subscribed := req.FormValue("subscribed")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<form method="get">
		<input type="text" name="fn" placeholder="first name">
		<input type="text" name="ln" placeholder="last name">
		<input type="checkbox" name="subscribed">
		<label>subscribe</label>
		<input type="submit">
	</form>
	<h1>First name: `+fn+`</h1>
	<h1>Last name: `+ln+`</h1>
	<h1>Subscribed? `+subscribed+`</h1>`)
}
