package main

import (
	"errors"
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
		return false, errors.New("your number is not positive")
	}

	return true, nil
}
