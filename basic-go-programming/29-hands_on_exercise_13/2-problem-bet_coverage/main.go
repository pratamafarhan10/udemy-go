package main

import (
	"fmt"

	"github.com/udemy-go/29-hands_on_exercise_13/2-problem-bet_coverage/quote"
	"github.com/udemy-go/29-hands_on_exercise_13/2-problem-bet_coverage/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
