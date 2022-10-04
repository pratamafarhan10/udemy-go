package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := `The dash, it's digi', the schedule busy
	My head in a hoodie, my shorty a goodie
	My cousins are crazy, my cousins like Boogie`

	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s64)

	d, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("error:", err)
	}

	fmt.Println(string(d))
}
