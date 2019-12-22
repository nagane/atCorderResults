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
	for i := min; ; {
		fmt.Print("i = ")
		fmt.Println(i)
		if i%min == 0 {
			if i > max {
				if i%max == 0 {
					fmt.Println(i)
					break
				}
			} else {
				if max%i == 0 {
					fmt.Println(i)
					break
				}
			}
		}
		i++
	}
}
