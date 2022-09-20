package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type Person struct {
	Name, Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	name := Person{
		Name:  "drake",
		Motto: "You only live once",
	}

	err := tpl.Execute(os.Stdout, name)
	if err != nil {
		log.Println(err)
	}
}
