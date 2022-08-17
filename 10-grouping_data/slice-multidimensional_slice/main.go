package main

import "fmt"

func main() {
	jamesBond := []string{"James", "Bond"}

	fmt.Println(jamesBond)

	vanoss := []string{"vanoss", "gaming"}

	fmt.Println(vanoss)

	people := [][]string{jamesBond, vanoss}

	fmt.Println(people)

	farhan := []string{"farhan", "pratama"}

	people = append(people, farhan)

	fmt.Println(people)

	kembar := [][]string{{"adul", "zulkifli"}, {"komeng", "syahputra"}}

	people = append(people, kembar...)

	fmt.Println(people)
}
