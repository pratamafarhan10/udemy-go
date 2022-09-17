package main

import (
	"fmt"
	"strings"

	"github.com/udemy-go/28-testing_and_benchmarking/benchmarking_examples/mystr"
)

func main() {
	story := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aenean eu turpis neque. Praesent non est non nulla posuere faucibus sit amet et nunc. Donec ipsum nibh, rhoncus in dapibus eu, pharetra nec mi. Donec facilisis volutpat mi at tempus. Nunc mollis augue id pretium faucibus. Aenean sed libero interdum, ornare eros congue, dapibus tortor. Suspendisse blandit nulla et tempus rhoncus. Aliquam sem lacus, tristique ac eleifend vitae, ullamcorper sit amet libero. Nunc sed est eros. Donec ultricies sodales turpis, at imperdiet orci lobortis in."

	str := strings.Split(story, " ")

	for _, val := range str {
		fmt.Println(val)
	}

	fmt.Println("=========== Cat vs Join")
	fmt.Printf("\n%v\n", mystr.Cat(str))
	fmt.Printf("\n%v\n", mystr.Join(str))

}
