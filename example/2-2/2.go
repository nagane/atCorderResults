package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]int, n)
	t := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&t[i])
	}
	fmt.Println(s)
	fmt.Println(t)

}
