package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	log.Println("server listens at:", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}

		fmt.Fprintln(conn, "Hello there")
		fmt.Fprintln(conn, "Nice to meet you my g")

		conn.Close()
	}
}
