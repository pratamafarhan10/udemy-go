package main

import (
	"fmt"

	"github.com/udemy-go/29-hands_on_exercise_13/3-problem-table_tests/mymath"
)

func main() {
	xxi := gen()
	for _, v := range xxi {
		fmt.Println(v[1 : len(v)-1])
		fmt.Println(mymath.CenteredAvg(v))
	}
}

// gen returns a slice of test cases
func gen() [][]int {
	a := []int{1, 4, 6, 8, 100}
	b := []int{0, 8, 10, 1000}
	c := []int{9000, 4, 10, 8, 6, 12}
	d := []int{123, 744, 140, 200}
	e := [][]int{a, b, c, d}
	return e
}
