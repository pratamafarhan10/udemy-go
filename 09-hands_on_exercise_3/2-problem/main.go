package main

import "fmt"

func main() {
	a := 1
	for i := 65; i <= 90; i++ {
		fmt.Println(a)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
		a++
	}
}
