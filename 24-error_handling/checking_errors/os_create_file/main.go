package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("name.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	_, err = f.Write([]byte(`We
can
use
string
raw
literals
as
well
	`))

	if err != nil {
		fmt.Println(err)
		return
	}

	rf, err := os.ReadFile("name.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(rf))
}
