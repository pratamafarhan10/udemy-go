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
			line := scanner.Text()

			s := strings.Fields(line)

			method := s[0]
			url := s[1]

			fmt.Println(method, url)

			if method == "GET" {
				if url == "/" {
					welcome(conn, "Welcome to my website")
				} else if url == "/ronaldo" {
					ronaldo(conn, "RONALDO THE GOAT👑🐐")
				} else if url == "/darwinnunez" {
					darwinNunez(conn, "Liverpool in the mud😭")
				} else {
					notFound(conn, "URL not found")
				}
			}
		}
		i++
	}
}

func respond(conn net.Conn, message, httpCode string) {
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
    <h1>` + message + `</h1>
</body>
</html>
	`

	fmt.Fprintf(conn, "HTTP/1.1 %v OK\r\n", httpCode)
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func welcome(conn net.Conn, message string) {
	respond(conn, message, "200")
}

func ronaldo(conn net.Conn, message string) {
	respond(conn, message, "200")
}

func darwinNunez(conn net.Conn, message string) {
	respond(conn, message, "200")
}

func notFound(conn net.Conn, message string) {
	respond(conn, message, "404")
}
