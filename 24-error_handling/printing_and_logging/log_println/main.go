package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.Create("log.txt")

	log.SetOutput(logFile)

	if err != nil {
		log.Println(err)
		return
	}

	defer logFile.Close()

	nf, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Println("err happened:", err)
		log.Println(err)
		return
	}

	defer nf.Close()

	nf2, err := os.Open("no-file.txt")

	if err != nil {
		fmt.Println("err happened:", err)
		log.Println(err)
		return
	}

	defer nf2.Close()
}
