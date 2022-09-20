package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"formatDate": formatDate,
}

func init() {
	// Must must receive *template and error
	// New return *template
	// Funcs return *template
	// ParseFiles return *template and error
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.html"))
}

func formatDate(s time.Time) string {
	return s.Format("02-01-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.html", time.Now())
	if err != nil {
		log.Println(err)
	}
}
