package main

import "fmt"

func main() {
	var A int
	var C [6]int
	V := [...]int{1, 5, 10, 50, 100, 500}
	fmt.Scan(&C[0], &C[1], &C[2], &C[3], &C[4], &C[5], &A)

	ans := 0

	for i := 5; i >= 0; i-- {
		var t int

		if A/V[i] > 0 && 0 < C[i] {
			if A/V[i] < C[i] {
				// コインの枚数
				t = A / V[i]
				A -= V[i] * (A / V[i])
			} else {
				// コインの枚数
				t = C[i]
				A -= V[i] * C[i]
			}

			ans += t
		}
	}

	fmt.Println(ans)
}
