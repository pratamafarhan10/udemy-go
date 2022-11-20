package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	n := "sule"

	changeValue(&n)

	fmt.Println(n)

	u := user{
		name: "sule",
		age:  42,
	}

	changeStruct(&u)
	fmt.Println(u)
}

func changeValue(name *string) {
	*name = "farhan"
}

// Kalo dikirim dengan tipe pointer, berarti langsung ubah sesuai memory addressnya
func changeStruct(user *user) {
	user.name = "farhan"
	user.age = 22
}
