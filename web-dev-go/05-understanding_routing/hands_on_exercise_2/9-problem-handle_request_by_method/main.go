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
	defer listener.Close()
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

func request(conn net.Conn) map[string]string {
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		s := strings.Fields(ln)
		fmt.Println(ln)

		if i == 0 {
			if s[0] == "GET" && s[1] == "/" {
				index(conn)
			} else if s[0] == "GET" && s[1] == "/apply" {
				applyGet(conn)
			} else if s[0] == "POST" && s[1] == "/apply" {
				applyPost(conn)
			}
		}

		if ln == "" {
			break
		}
		i++
	}
	return nil
}

func handle(conn net.Conn, i int) {
	fmt.Println(i)
	defer conn.Close()

	request(conn)
}

func index(conn net.Conn) {
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
    <h1>HELLO WORLD</h1>
</body>
</html>
	`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyGet(conn net.Conn) {
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
    <h1>Apply</h1>
	<form action="/apply" method="POST">
        <button type="submit">Hi</button>
    </form>
</body>
</html>
	`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func applyPost(conn net.Conn) {
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
    <h1>Apply Post</h1>
</body>
</html>
	`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
