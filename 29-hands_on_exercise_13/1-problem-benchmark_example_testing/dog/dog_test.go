package dog

import (
	"fmt"
	"testing"
)

type testPair struct {
	value, expected int
}

var testCases = []testPair{
	{2, 14},
	{1, 7},
	{3, 21},
}

func TestYears(t *testing.T) {
	for _, val := range testCases {
		v := Years(val.value)

		if v != val.expected {
			t.Error("Expected", val.expected, "Got", v)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	for _, val := range testCases {
		v := YearsTwo(val.value)

		if v != val.expected {
			t.Error("Expected", val.expected, "Got", v)
		}
	}
}

func ExampleYears() {
	fmt.Println(Years(2))
	// Output:
	// 14
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(2))
	// Output:
	// 14
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(2)
	}
}
