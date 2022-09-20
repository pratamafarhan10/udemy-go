package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")

	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create("index.html")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	err = tpl.Execute(f, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
