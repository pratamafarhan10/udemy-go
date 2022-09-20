package main

import (
	"fmt"
	"log"
)

type positiveNumberError struct { // this type implement type Error as well because it has Error method attached to it
	number     int
	isPositive bool
	err        error
}

func (pne positiveNumberError) Error() string {
	return fmt.Sprintf("Number is not positive: number (%v) isPositive (%v) errMessage (%v)", pne.number, pne.isPositive, pne.err)
}

func main() {
	pne1 := positiveNumberError{
		number:     -10,
		isPositive: false,
		err:        fmt.Errorf("yo this is an error"),
	}

	fmt.Println(pne1) // So basically if a type implement Error, it will run the Error method automatically

	res, err := positive(-1)

	if err != nil {
		fmt.Printf("The type of err: %T\n", err)
		log.Println(err)
		return
	}

	fmt.Println(res)
}

func positive(i int) (bool, error) {
	if i < 0 {
		message := fmt.Errorf("not positive")
		return false, positiveNumberError{i, false, message}
	}

	return true, nil
}
