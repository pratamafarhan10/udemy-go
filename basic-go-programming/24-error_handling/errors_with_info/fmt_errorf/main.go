package main

import (
	"fmt"
	"log"
)

func main() {
	res, err := positive(-1)

	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res)
}

func positive(i int) (bool, error) {
	if i < 0 {
		return false, fmt.Errorf("your number is %v, and is not positive", i)
	}

	return true, nil
}
