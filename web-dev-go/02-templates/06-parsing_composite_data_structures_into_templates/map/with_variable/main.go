package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	names := map[string]string{
		"rap":    "drake",
		"hiphop": "chris brown",
		"actor":  "james bond",
	}

	err := tpl.Execute(os.Stdout, names)
	if err != nil {
		log.Println(err)
	}
}
