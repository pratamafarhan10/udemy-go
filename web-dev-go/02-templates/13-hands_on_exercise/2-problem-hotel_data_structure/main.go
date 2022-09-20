package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Region string
	Zip                         int
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	californiaHotels := []hotel{
		{
			Name:    "Grand Hyatt",
			Address: "Jl. California",
			City:    "Los Angeles",
			Region:  "Northern",
			Zip:     189023,
		},
		{
			Name:    "Ibis",
			Address: "Jl. Drake",
			City:    "Downton Los Angeles",
			Region:  "Southern",
			Zip:     192345,
		},
	}

	err := tpl.Execute(os.Stdout, californiaHotels)
	if err != nil {
		log.Println(err)
	}
}
