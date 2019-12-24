package main

import (
	"fmt"
)

func main() {
	// 100000
	var A, B int
	fmt.Scan(&A, &B)
	var min, max int
	if A < B {
		min = A
		max = B
	} else {
		min = B
		max = A
	}

	for i := 1; ; {
		fmt.Print("i = ")
		fmt.Println(i)
		if (i*min)%max == 0 {
			fmt.Println(i * min)
			break
		}
		i++
	}
}
