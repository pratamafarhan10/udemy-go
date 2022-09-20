package greet

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	v := Greet("Farhan")

	if v != "Welcome my dear Farhan" {
		t.Error("Expected:", "Welcome my dear Farhan", "Got:", v)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Farhan"))
	// Output:
	// Welcome my dear Farhan
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Farhan")
	}
}
