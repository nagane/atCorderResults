package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)
	// A x N + B x d(N)
	buyNum := 1000000000
	buyedNumber(buyNum, a, b, x)
	fmt.Println(strconv.Itoa(buyNum))

}

// n 買う数字
// a 係数
// x 所持金
func buyedNumber(n int, a int, b int, x int) bool {

	price := a*n + b*len(strconv.Itoa(n))
	if x >= price {
		buyedNumber(n/2, a, b, x)
	}
	return true
}
