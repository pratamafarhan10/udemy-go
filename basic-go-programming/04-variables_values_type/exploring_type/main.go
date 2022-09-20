package main

import "fmt"

var x = 42

var y = "farhan" // golang is static programming language so if a variable assigned as a string, you can't change the data type

var z = `he called me 

"farhan"

`

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(y)
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
