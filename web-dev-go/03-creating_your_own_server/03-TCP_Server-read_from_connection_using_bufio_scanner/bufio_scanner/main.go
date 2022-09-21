package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "The dash, it's digital, the schedule busy\n My head in a hoodie, my shorty a goodie\n My cousins are crazy, my cousins like Boogie\n Life is amazin', it is what it should be\n Been here for ten but I feel like a rookie"

	scanner := bufio.NewScanner(strings.NewReader(s))

	// Scan per byte
	// scanner.Split(bufio.ScanBytes)

	// Scan per line
	// scanner.Split(bufio.ScanLines)

	// Scan per rune returns each UTF-8-encoded rune as a token
	// scanner.Split(bufio.ScanRunes)

	// Scan() will return by default per line
	for scanner.Scan() {
		b := scanner.Text()
		fmt.Println(b)
	}
}
