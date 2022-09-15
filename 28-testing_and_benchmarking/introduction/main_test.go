package main

import "testing"

type testPair struct {
	value, expected int
}

func TestFactorial(t *testing.T) {
	testCases := []testPair{
		{4, 24},
		{5, 120},
		{10, 3628800},
	}

	for _, val := range testCases {
		v := factorial(val.value)

		if v != val.expected {
			t.Error("Expected", val.expected, "Got", v)
		}
	}

}
