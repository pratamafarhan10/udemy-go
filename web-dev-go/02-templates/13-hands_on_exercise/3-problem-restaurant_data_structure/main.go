package main

import (
	"log"
	"os"
	"text/template"
)

type menu struct {
	Occasion, Name string
	Price          int
}

type restaurant struct {
	Name, Location string
	Menus          []menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	restaurants := []restaurant{
		{
			Name:     "McDonalds",
			Location: "Gaplek",
			Menus: []menu{
				{"Breakfast", "Happy Meal", 20000},
				{"Lunch", "Big Mac", 40000},
				{"Dinner", "Spicy Wings", 30000},
			},
		},
		{
			Name:     "Warteg",
			Location: "Sukabirus",
			Menus: []menu{
				{"Breakfast", "Bubur Ayam", 10000},
				{"Lunch", "Nasi Goreng", 12000},
				{"Dinner", "Pecel Lele", 15000},
			},
		},
	}

	err := tpl.Execute(os.Stdout, restaurants)
	if err != nil {
		log.Println(err)
	}
}
