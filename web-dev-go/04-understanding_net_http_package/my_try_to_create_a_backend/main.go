package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type clientHandler struct{}

func main() {
	var ch clientHandler
	http.ListenAndServe(":8080", ch)
}

func (c clientHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err.Error())
	}

	data := struct {
		Method        string
		URL           *url.URL
		Submissions   map[string][]string
		Header        map[string][]string
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	b, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Fprintln(w, string(b))

}
