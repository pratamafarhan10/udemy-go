package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Server listens at 8080")
	defer listener.Close()

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Fatalln(err.Error())
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		if i == 0 {
			ln := scanner.Text()

			s := strings.Fields(ln)[1]

			fmt.Println(s)
			respond(conn, s)
		}
		i++
	}
	defer conn.Close()
}

func respond(conn net.Conn, url string) {
	body := `
	<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>` + url + `</h1>
</body>
</html>
	`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
