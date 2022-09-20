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
	names := []Person{
		{"drake", "You only live once"},
		{"lil baby", "I always got the same process. Now I can flow a little better"},
		{"eminem", "I got a list, here's the order of my list that it's in/It goes Reggie, Jay Z, 2Pac, and Biggie, Andre from Outkast, Jada, Kurupt, Nas, and then me"},
	}

	err := tpl.Execute(os.Stdout, names)
	if err != nil {
		log.Println(err)
	}
}
