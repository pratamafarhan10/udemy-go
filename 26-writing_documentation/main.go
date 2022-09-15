package main

import (
	"fmt"

	"github.com/udemy-go/26-writing_documentation/mymath"
)

func main() {
	fmt.Println("1 + 1:", mymath.Sum(1, 1))
	fmt.Println("1 + 3:", mymath.Sum(1, 3))
	fmt.Println("10 + 4:", mymath.Sum(10, 4))
}
