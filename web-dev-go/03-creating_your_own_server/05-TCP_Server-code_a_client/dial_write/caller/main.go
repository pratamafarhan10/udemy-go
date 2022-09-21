package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintln(conn, "Hello from caller")
	fmt.Fprintln(conn, "Wagwan my g")

	conn.Close()
}
