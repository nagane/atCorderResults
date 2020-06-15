package main

import "fmt"

var x, y int

func main() {
	fmt.Scan(&x, &y)

	ans := "No"
	// a = 1だと多分全部4で通るみたいなケースがダメっぽい？
	for a := 0; a < x+1; a++ {
		b := x - a
		if 2*a+4*b == y {
			ans = "Yes"
		}
	}

	fmt.Println(ans)

}
