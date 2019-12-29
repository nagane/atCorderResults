package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	for i := x; ; i++ {
		if x%i == 0 {
			if i >= x {
				fmt.Println(i)
				break
			}
		}
	}
}
