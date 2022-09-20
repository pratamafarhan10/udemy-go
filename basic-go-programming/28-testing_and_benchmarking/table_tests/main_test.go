package main

import "testing"

type testPair struct {
	numbers  []int
	expected int
}

// Function TestMySum is to test the function MySum in package main
func TestMySum(t *testing.T) {
	testCases := []testPair{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 21}, 22},
		{[]int{21, 21}, 42},
	}

	for _, val := range testCases {
		g := mySum(val.numbers...)

		if g != val.expected {
			t.Error("Expected", val.expected, "Got", g)
		}
	}
}
