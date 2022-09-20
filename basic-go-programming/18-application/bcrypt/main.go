package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	fmt.Println("==== Generate encrypted password")
	pass := `password123`

	enc, err := bcrypt.GenerateFromPassword([]byte(pass), 10)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(pass)
	fmt.Println(string(enc))

	fmt.Println("==== Check password")
	loginPass := `password1243`

	err = bcrypt.CompareHashAndPassword(enc, []byte(loginPass))

	if err != nil {
		fmt.Println("Wrong password")
		return
	}
	fmt.Println("Correct password")
}
