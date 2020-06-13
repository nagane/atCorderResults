package main

import "fmt"

func main() {
	for i := 1; i < 6; i++ {
		var x int
		fmt.Scan(&x)
		if x == 0 {
			fmt.Println(i)
		}
	}
}
