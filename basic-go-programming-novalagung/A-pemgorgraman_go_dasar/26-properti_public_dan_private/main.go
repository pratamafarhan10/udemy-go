package main

import (
	"fmt"

	"github.com/udemy-go/basic-go-programming-novalagung/A-pemgorgraman_go_dasar/26-properti_public_dan_private/library"
	// . "github.com/udemy-go/basic-go-programming-novalagung/A26-properti_public_dan_private/library" // menjadikan package library satu level dengan package di main
	// lib "github.com/udemy-go/basic-go-programming-novalagung/A26-properti_public_dan_private/library" // package alias, jadi dipanggilnya bukan library tapi lib
)

// import "/library"

func main() {
	fmt.Println("============= File di dalam package yang berbeda")
	library.SayHello("ethan")
	// library.introduce("ethan")
	s1 := library.Student{
		Name:  "Andre",
		Grade: 82,
	}

	fmt.Println(s1.Name, s1.Grade)

	fmt.Println("============= File di dalam package yang sama")
	sayHello("farhan")
	// * Fungsi sayHello pada file partial.go bisa dikenali meski level aksesnya adalah unexported. Hal ini karena kedua file tersebut (main.go dan partial.go) memiliki package yang sama
	// * Cara lain untuk menjalankan program bisa dengan perintah go run *.go, dengan cara ini tidak perlu menuliskan nama file nya satu per satu.

	fmt.Println("============= Func init di library package")
	fmt.Println(library.Student1)
}
