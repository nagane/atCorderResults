package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	cnt := 0
	e := len(S) - 1
	for i := 0; i < len(S)/2; i++ {
		if S[i] != S[e] {
			cnt++
		}
		e--
	}
	fmt.Print(cnt)
}
