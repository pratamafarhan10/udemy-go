package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	b, err := io.ReadAll(conn)

	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(b))
}
