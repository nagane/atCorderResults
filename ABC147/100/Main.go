package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if 21 < (a + b + c) {
		fmt.Print("bust")
	} else {
		fmt.Print("win")
	}
	return
}
