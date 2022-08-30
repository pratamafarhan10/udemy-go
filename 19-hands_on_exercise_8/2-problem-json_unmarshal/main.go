package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {
	request := `[{"Name":"Andre","Age":42},{"Name":"Sule","Age":39},{"Name":"Komeng","Age":51}]`

	var people []Person

	err := json.Unmarshal([]byte(request), &people)

	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Println(people)
}
