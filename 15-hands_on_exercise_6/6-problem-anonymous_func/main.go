package main

import "fmt"

func main() {
	func() {
		fmt.Println("this is a anonymous function")
	}()
}
