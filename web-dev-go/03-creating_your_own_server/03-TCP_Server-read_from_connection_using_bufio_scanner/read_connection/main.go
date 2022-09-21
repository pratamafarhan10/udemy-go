package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	log.Println("Server listening at 8080")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		b := scanner.Text()
		fmt.Println(b)
	}
	defer conn.Close()

	// We never got here
	// we have an open stream connection
	// how does the above reader know when it's done?
	fmt.Println("Code got here.")
}
