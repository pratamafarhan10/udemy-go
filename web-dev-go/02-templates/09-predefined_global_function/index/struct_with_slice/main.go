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
	data := struct {
		Rappers []string
		Cars    []string
	}{
		Rappers: []string{"drake", "lil baby", "eminem"},
		Cars:    []string{"toyota", "honda", "hyundai"},
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Println(err)
	}
}
