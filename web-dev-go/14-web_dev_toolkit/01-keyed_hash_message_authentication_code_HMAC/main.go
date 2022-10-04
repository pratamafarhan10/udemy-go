package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

func main() {
	a := getCode("windah")
	fmt.Println(a)
	b := getCode("basudara")
	fmt.Println(b)

	c := hmac.Equal([]byte(a), []byte(b))
	fmt.Println(c)
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, s)
	return fmt.Sprintf("%x", string(h.Sum(nil)))
}
