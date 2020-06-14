package main

import "fmt"

var x, y int

func main() {
	fmt.Scan(&x, &y)

	var flag bool
	sum := y / 4
	s := y / 4
	iter := 0

	for {
		if sum > x {
			flag = false
			break
		}
		if (s*4)+(iter*2) > y {
			flag = false
			break
		}
		if sum == x {
			flag = true
			break
		}
		sum++
		iter++
		s--
	}

	if flag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
