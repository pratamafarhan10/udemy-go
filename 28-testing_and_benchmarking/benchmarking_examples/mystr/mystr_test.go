package mystr

import (
	"fmt"
	"strings"
	"testing"
)

func TestCat(t *testing.T) {
	v := Cat([]string{"Andre", "Taulany"})

	if v != "Andre Taulany" {
		t.Error("Expected: Andre Taulany. Got:", v)
	}
}

func TestJoin(t *testing.T) {
	v := Join([]string{"Andre", "Taulany"})

	if v != "Andre Taulany" {
		t.Error("Expected: Andre Taulany. Got:", v)
	}
}

func ExampleCat() {
	fmt.Println(Cat([]string{"Andre", "Taulany"}))
	// Output:
	// Andre Taulany
}

func ExampleJoin() {
	fmt.Println(Join([]string{"Andre", "Taulany"}))
	// Output:
	// Andre Taulany
}

func BenchmarkCat(b *testing.B) {
	story := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean eu turpis neque. Praesent non est non nulla posuere faucibus sit amet et nunc. Donec ipsum nibh, rhoncus in dapibus eu, pharetra nec mi. Donec facilisis volutpat mi at tempus. Nunc mollis augue id pretium faucibus. Aenean sed libero interdum, ornare eros congue, dapibus tortor. Suspendisse blandit nulla et tempus rhoncus. Aliquam sem lacus, tristique ac eleifend vitae, ullamcorper sit amet libero. Nunc sed est eros. Donec ultricies sodales turpis, at imperdiet orci lobortis in."

	str := strings.Split(story, " ")
	for i := 0; i < b.N; i++ {
		Cat(str)
	}
}

func BenchmarkJoin(b *testing.B) {
	story := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean eu turpis neque. Praesent non est non nulla posuere faucibus sit amet et nunc. Donec ipsum nibh, rhoncus in dapibus eu, pharetra nec mi. Donec facilisis volutpat mi at tempus. Nunc mollis augue id pretium faucibus. Aenean sed libero interdum, ornare eros congue, dapibus tortor. Suspendisse blandit nulla et tempus rhoncus. Aliquam sem lacus, tristique ac eleifend vitae, ullamcorper sit amet libero. Nunc sed est eros. Donec ultricies sodales turpis, at imperdiet orci lobortis in."

	str := strings.Split(story, " ")
	for i := 0; i < b.N; i++ {
		Join(str)
	}
}
