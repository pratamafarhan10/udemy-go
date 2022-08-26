package main

import "fmt"

func main() {
	bornYear := 2000

	for {
		if bornYear > 2022 {
			break
		}
		fmt.Println(bornYear)
		bornYear++
	}
}
