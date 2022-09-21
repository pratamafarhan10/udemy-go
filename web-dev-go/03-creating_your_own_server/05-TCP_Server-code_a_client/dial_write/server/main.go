package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}
	defer listener.Close()

	log.Println("Server listens at:", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()

		fmt.Println(ln)
		// fmt.Fprintln(conn, ln)
	}

	defer conn.Close()
}
