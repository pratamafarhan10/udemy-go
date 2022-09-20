package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	name := struct {
		Name, Country string
	}{
		"James Bond",
		"England",
	}
	err := tpl.Execute(os.Stdout, name)
	if err != nil {
		log.Println(err)
	}
}
