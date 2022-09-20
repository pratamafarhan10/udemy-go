package main

import (
	"fmt"
	"os"
)

func main() {
	name := "Todd Mcleod"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World! </title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`

	f, err := os.Create("index.html")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	_, err = f.Write([]byte(tpl))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tpl)
}
