package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "Andre",
		Age:  42,
	}

	p2 := Person{
		Name: "Sule",
		Age:  39,
	}

	p3 := Person{
		Name: "Komeng",
		Age:  51,
	}

	persons := []Person{p1, p2, p3}

	fmt.Println(persons)

	res, err := json.Marshal(persons)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(res))
}
