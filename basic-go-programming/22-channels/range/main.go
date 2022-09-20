package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 1; i < 6; i++ {
			c <- i
		}

		close(c) // Menutup channel ketika sudah selesai, kalau tidak ditutup maka terjadi error (deadlock)
	}()

	for val := range c {
		fmt.Println(val)
	}

	fmt.Println("Exit")
}
