package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	log.Println("Server running at 8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		io.WriteString(conn, "Hello world\n")
		fmt.Fprint(conn, "this is from fprint\n")
		fmt.Fprintf(conn, "%v", "this is from fprintf \n")
		fmt.Fprintln(conn, "this is from fprintln")
		conn.Close()
	}
}
