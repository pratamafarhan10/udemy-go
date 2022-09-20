package main

import "testing"

type testPair struct {
	value, expected int
}

var testCases = []testPair{
	{5, 120},
	{4, 24},
	{3, 6},
}

func TestFactorial(t *testing.T) {
	for _, val := range testCases {
		v := factorial(val.value)
		if v != val.expected {
			t.Errorf(`
			Error!
			Value: %v
			Expected: %v
			Got: %v
			`, val.value, val.expected, v)
		}
	}
	// v := factorial(4)

	// if v != 24 {
	// 	t.Error("Expected 24 got", v)
	// }
}
