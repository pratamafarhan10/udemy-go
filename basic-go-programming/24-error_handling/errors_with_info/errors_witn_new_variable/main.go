package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrorMessage = errors.New("your number is not positive")

func main() {
	fmt.Printf("error message type: %T\n", ErrorMessage)
	res, err := positive(-1)

	if err != nil {
		log.Fatalln(err)
		return
	}

	fmt.Println(res)
}

func positive(i int) (bool, error) {
	if i < 0 {
		return false, ErrorMessage
	}

	return true, nil
}
