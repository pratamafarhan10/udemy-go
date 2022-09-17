package mymath

import (
	"fmt"
	"testing"
)

type testPair struct {
	value    []int
	expected float64
}

var testCases = []testPair{
	{[]int{1, 4, 6, 8, 100}, 6},
	{[]int{0, 8, 10, 1000}, 9},
	{[]int{9000, 4, 10, 8, 6, 12}, 9},
	{[]int{123, 744, 140, 200}, 170},
}

func TestCenteredAvg(t *testing.T) {
	for _, testCase := range testCases {
		v := CenteredAvg(testCase.value)

		if v != testCase.expected {
			t.Error("Expected", testCase.expected, "Got", v)
		}
	}
}

func ExampleCenteredAvg() {
	fmt.Println(CenteredAvg([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}

func BenchmarkCenteredAvg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CenteredAvg([]int{1, 4, 6, 8, 100})
	}
}
