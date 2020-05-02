package main

import (
	"fmt"
	"math"
)

func main() {
	var X int
	fmt.Scan(&X)

	var m float64
	m = 100.0

	for i := 1; ; i++ {
		m = math.Trunc(m * 1.01)

		if int(m) >= X {
			fmt.Println(i)
			break
		}
	}
}
