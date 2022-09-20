package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{53, 35, 1, 64, 2, 63, 3, 66}
	words := []string{"jimmy", "kevin", "adul", "sule", "desta"}

	fmt.Println("before sort:", numbers)
	sort.Ints(numbers)
	fmt.Println("after sort:", numbers)
	ias := sort.IntsAreSorted(numbers)
	fmt.Println("ints are sorted:", ias)
	searchInt := sort.SearchInts(numbers, 2)
	fmt.Println("search the index of 2:", searchInt)

	fmt.Println("----------")
	fmt.Println("before sort:", words)
	sort.Strings(words)
	fmt.Println("after sort:", words)
	sas := sort.StringsAreSorted(words)
	fmt.Println("string are sorted:", sas)
	searchStr := sort.SearchStrings(words, "kevin")
	fmt.Println("search the index of kevin:", searchStr)

	fmt.Println("---------- Search guessing game")
	// GuessingGame()

	fmt.Println("---------- Scanf")
	var ab string
	var ac string
	fmt.Scanf("%v %v", &ab, &ac)
	fmt.Println(ab, ac)
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)

		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
