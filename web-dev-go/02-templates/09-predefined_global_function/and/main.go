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

type rapper struct {
	Name, Motto string
	Goat        bool
	IsActive    bool
}

func main() {
	rappers := []rapper{
		{"drake", "You only live once", true, true},
		{"lil baby", "I always got the same process. Now I can flow a little better", false, true},
		{"eminem", "I got a list, here's the order of my list that it's in/It goes Reggie, Jay Z, 2Pac, and Biggie, Andre from Outkast, Jada, Kurupt, Nas, and then me", true, false},
	}

	err := tpl.Execute(os.Stdout, rappers)
	if err != nil {
		log.Println(err)
	}
}
