package main

import "fmt"

func main() {
	var c1, c5, c10, c50, c100, c500, A int
	var C [5]int
	V := [...]int{1, 5, 10, 50, 100, 500}
	fmt.Scan(&C[0], &C[1], &C[2], &C[3], &C[4], &C[5], A)

	ans := 0

	for i := 5; i >= 0; i-- {
		var t int

		if A/V[i] > C[i] {

		}

	}
}
