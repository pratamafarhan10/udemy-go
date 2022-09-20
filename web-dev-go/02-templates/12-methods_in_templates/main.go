package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

func (p person) DoubleAge() int {
	return p.Age * 2
}

func (p person) MultiplyAge(age int) int {
	return age * 2
}

func (p person) Introduction() string {
	return fmt.Sprint("wagwan my g, man's name is ", p.Name)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	jj := person{
		Name: "JJ",
		Age:  29,
	}

	err := tpl.Execute(os.Stdout, jj)
	if err != nil {
		log.Println(err)
	}
}
