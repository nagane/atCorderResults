package main

import "fmt"

func main() {
	var K int
	fmt.Scan(&k)
	ans := 0
	for a := 1; a <= K; a++ {
		for b := 1; b <= K; b++ {
			g := gcd(a, b)
			for c := 1; c <= K; c++ {
				ans += gcd(g, c)
			}
		}
	}
}

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}
