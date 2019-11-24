package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)

	for i := 0; i < len(s); i++ {
		modifyChar(n, s[i])
	}
}

func modifyChar(n int, s byte) {
	// A = 65 Z = 90
	c := n + int(s)
	if c > 90 {
		c = c - 90 + 64
	}
	fmt.Print(string(c))
}
