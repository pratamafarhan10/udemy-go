package main

import "fmt"

var x = 42

// var y = 20

func main() {
	// Println = buat cetak dan enter
	fmt.Println("Println:", x)
	// Print = buat cetak tanpa enter
	fmt.Print("Print:", x, "\n")
	// Printf = buat cetak format
	fmt.Printf("%T\n", x)

	// Sprintln = mengembalikan string dan enter
	a := fmt.Sprintln(x)
	fmt.Print("Sprintln:", a)
	// Sprint = mengembalikan string tanpPrintln()
	b := fmt.Sprint(x)
	fmt.Println("Sprint:", b)
	// Sprintf = mengembalikan format dalam bentuk string
	c := fmt.Sprintf("%T", x)
	fmt.Println("Sprintf:", c)

	if fmt.Sprintf("%T", x) == "int" {
		fmt.Println("Hai ini masuk ke if")
	}

	// fmt.Println(x)
	// fmt.Printf("%T\n", x)  // int
	// fmt.Printf("%#v\n", x) // 42
	// fmt.Printf("%x\n", x)  //2a
	// a, _ := fmt.Printf("%b%X%x\n", x, x, x)
	// fmt.Println("a:", a)

	// s := fmt.Sprintf("%b\t%X\t%x", x, x, x)
	// fmt.Println("println s", s)
	// fmt.Printf("%T\n", a)
	// fmt.Printf("%T\n", s)

}
