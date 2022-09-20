package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)

type animal struct {
	Name, Order string
}

func main() {
	fmt.Println("============== marshal")
	reds := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Ruby", "Maroon"},
	}

	fmt.Println("before marshal", reds)

	result, err := json.Marshal(reds)

	if err != nil {
		fmt.Println("error:", err)
	}

	os.Stdout.Write(result)
	fmt.Printf("\nafter marshal: %v\n", result)

	convertedSliceOfByte := string(result[:])

	fmt.Println(convertedSliceOfByte)

	fmt.Println("============== Unmarshal")

	var animals []animal
	err2 := json.Unmarshal(jsonBlob, &animals)

	if err2 != nil {
		fmt.Println("error:", err2)
	}

	fmt.Println(animals)
}
