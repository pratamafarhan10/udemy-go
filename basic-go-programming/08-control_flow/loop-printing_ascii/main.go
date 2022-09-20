package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("Number: %v - String: %#U\n", i, i)
	}
}
