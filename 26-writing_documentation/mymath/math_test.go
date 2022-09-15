package mymath

import "testing"

type testPair struct {
	value1, value2, expected int
}

var testCases = []testPair{
	{1, 2, 3},
	{2, 2, 4},
	{2, 4, 6},
}

func TestSum(t *testing.T) {
	for _, val := range testCases {
		v := Sum(val.value1, val.value2)

		if v != val.expected {
			t.Errorf("Failed! value 1: %v, value 2: %v, expected: %v, got: %v", val.value1, val.value2, val.expected, v)
		}
	}
}
