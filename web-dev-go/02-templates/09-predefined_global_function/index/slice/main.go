package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.html"))
}

func main() {
	names := []string{"james", "bond", "drake", "lil baby"}

	err := tpl.Execute(os.Stdout, names)
	if err != nil {
		log.Println(err)
	}
}
