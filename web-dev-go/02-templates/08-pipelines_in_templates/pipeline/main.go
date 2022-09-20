package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"double": doubleNumber,
	"square": squareNumber,
	"sqrt":   squareRootNumber,
}

func init() {
	// Must must receive *template and error
	// New return *template
	// Funcs return *template
	// ParseFiles return *template and error
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.html"))
}
func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", 3.0)
	if err != nil {
		log.Println(err)
	}
}

func doubleNumber(n float64) float64 {
	return n + n
}

func squareNumber(n float64) float64 {
	return math.Pow(n, 2)
}

func squareRootNumber(n float64) float64 {
	return math.Sqrt(n)
}
