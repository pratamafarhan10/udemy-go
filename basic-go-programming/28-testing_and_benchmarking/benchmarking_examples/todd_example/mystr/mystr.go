package mystr

import "strings"

func Cat(c []string) string {
	s := c[0]

	for _, val := range c[1:] {
		s += " "
		s += val
	}

	return s
}

func Join(c []string) string {
	return strings.Join(c, " ")
}
