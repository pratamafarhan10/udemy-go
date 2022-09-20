// Package word provides calculation for words
package word

import "strings"

//  UseCount returns the counter of every word
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Count returns the length of a string
func Count(s string) int {
	// write the code for this func
	return len(s)
}
