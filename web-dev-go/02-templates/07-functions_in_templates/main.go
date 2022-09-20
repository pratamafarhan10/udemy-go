package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

func init() {
	// Must must receive *template and error
	// New return *template
	// Funcs return *template
	// ParseFiles return *template and error
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.html"))
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", data)
	if err != nil {
		log.Println(err)
	}
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
}
