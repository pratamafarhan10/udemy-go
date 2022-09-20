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
}

type car struct {
	Manufacturer, Type string
}

func main() {
	rappers := []rapper{
		{"drake", "You only live once"},
		{"lil baby", "I always got the same process. Now I can flow a little better"},
		{"eminem", "I got a list, here's the order of my list that it's in/It goes Reggie, Jay Z, 2Pac, and Biggie, Andre from Outkast, Jada, Kurupt, Nas, and then me"},
	}

	cars := []car{
		{"Toyota", "Innova"},
		{"Ferrari", "F40"},
		{"Lambhorgini", "Aventador"},
	}

	data := struct {
		Rappers []rapper
		Cars    []car
	}{
		Rappers: rappers,
		Cars:    cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Println(err)
	}
}
