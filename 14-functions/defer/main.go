package main

import "fmt"

// Defer digunakan untuk mengakhirkan eksekusi sebuah statement tepat sebelum blok fungsi selesai

func main() {
	defer foo()
	bar()

	a := 1
	b := 2

	defer sum(a, b) // bakal dijalanin pas barisnya dipanggil tapi dieksekusi pas udah selesai functionnya

	a++

	sum(a, b)

	fmt.Println(f())

	fmt.Println("============")
	orderSomeFood("pizza")
	orderSomeFood("burger")
	fmt.Println("============")

	fmt.Println(bb(2))
}

func sum(a, b int) {
	fmt.Println("Result:", a+b)
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func f() (result int) {
	fmt.Println(result)
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}

func bb(number int) int {
	defer func() {
		number *= 10
	}()

	return 2 * number
}

func orderSomeFood(food string) {
	defer fmt.Println("Terima kasih, silahkan tunggu")

	if food == "pizza" {
		fmt.Println("Pilihan yang tepat!")
		fmt.Println("Pizza disini yang paling enak!")
		return
	}

	fmt.Println("Pesanan anda:", food)
}
