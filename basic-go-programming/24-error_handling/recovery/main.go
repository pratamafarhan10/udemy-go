package main

import "fmt"

func main() {
	f()
	fmt.Println("Recovering from f")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from G: ", r)
		}
	}()
	fmt.Println("Calling G")
	g(0) // Tidak menjalankan baris dibawahnya karena function G panic. Langsung menjalankan defer
	fmt.Println("Recovering from G")
}

func g(i int) {
	if i > 3 {
		fmt.Println("PANICKING!!!")
		panic(fmt.Sprintf("panic: %v", i))
	}
	defer fmt.Println("Defer in G:", i)
	fmt.Println("Printing in G:", i)
	g(i + 1)
}
