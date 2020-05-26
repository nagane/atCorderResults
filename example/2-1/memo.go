package main

import (
	"fmt"
)

var memo []int

func main() {
	var n int

	fmt.Scan(&n)

	memo = make([]int, n+1)

	fmt.Println(fib(n))
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = fib(n-1) + fib(n-2)

	return memo[n]
}
