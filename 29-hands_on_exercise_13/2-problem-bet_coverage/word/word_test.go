package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	type testPair struct {
		value    string
		expected int
	}

	testCases := []testPair{
		{"chris brown", 11},
		{"drake", 5},
		{"sule", 4},
	}

	for _, val := range testCases {
		v := Count(val.value)

		if v != val.expected {
			t.Error("Expected", val.expected, "Got", v)
		}
	}
}

func TestUseCount(t *testing.T) {
	type testPair struct {
		value    string
		expected map[string]int
	}

	testCases := []testPair{
		{"chris brown", map[string]int{"chris": 1, "brown": 1}},
		{"drake", map[string]int{"drake": 1}},
		{"andre andre taulany", map[string]int{"andre": 2, "taulany": 1}},
	}

	for _, test := range testCases {
		words := UseCount(test.value)

		for k := range test.expected {
			if test.expected[k] != words[k] {
				t.Error("Expected", test.expected[k], "got", words[k])
			}
		}

	}

	// for _, val := range testCases {
	// 	words := UseCount(val.value)
	// 	split := strings.Fields(val.value)

	// 	for _, e := range split {
	// 		if words[e] != val.expected[e] {
	// 			t.Error("Expected", val.expected[e], "got", words[e])
	// 		}
	// 	}
	// }

}

func ExampleCount() {
	fmt.Println(Count("drake"))
	// Output:
	// 5
}

func ExampleUseCount() {
	words := UseCount("Chris Brown")

	for _, word := range words {
		fmt.Println(word)
	}

	// Output:
	// 1
	// 1
}

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("drake")
	}
}
