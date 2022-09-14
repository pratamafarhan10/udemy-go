package main

import "fmt"

type customErr struct {
	err error
}

func (c customErr) Error() string {
	return fmt.Sprintf("Error message: %v", c.err)
}

func main() {
	e1 := customErr{
		err: fmt.Errorf("this is an error message"),
	}

	foo(e1)
}

func foo(errMsg error) {
	fmt.Println(errMsg)
}
