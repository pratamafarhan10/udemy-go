package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err.Error())
	}
	i := 0
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalln(err.Error())
		}

		go handle(conn, i)
		i++
	}
}

func handle(conn net.Conn, i int) {
	fmt.Println(i)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()

	// The lines below is never written because the code above will keep scanning.
	// As long as we accept the request (or in other words we accept the call)
	// the program will read (scan) over and over until the code is over.
	// If there is nothing to scan for (for example the user close the web page)
	// then the scanning is done and it will run the code below
	fmt.Println("Code got here")
	io.WriteString(conn, "I see you connected")
}
