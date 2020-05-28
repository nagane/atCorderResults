package main

import (
	"fmt"
)

// DFS : Depth-Fisrt Search

//ついに禁忌をおかしたかもしれん
var a []int
var n, k int

func main() {
	fmt.Scan(&n)

	a = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Sacn(&k)

	calc(9, 0)
}

func calc(sum int, iter int) bool {

}
