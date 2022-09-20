package dog

import "testing"

type testPair struct {
	value, expected int
}

var testCases = []testPair{
	{1, 7},
	{2, 14},
	{3, 21},
	{4, 28},
}

func TestYears(t *testing.T) {
	for _, val := range testCases {
		v := Years(val.value)
		if v != val.expected {
			t.Errorf("Failed! value: %v, expected: %v, got: %v", val.value, val.expected, v)
		}
	}
}
