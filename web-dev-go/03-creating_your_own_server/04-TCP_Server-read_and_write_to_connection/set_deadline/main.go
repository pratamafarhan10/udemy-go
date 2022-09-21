package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	log.Println("Server listens at 8080")

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("TIMEOUT!")
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %v\n", ln)
	}

	defer conn.Close()
}
