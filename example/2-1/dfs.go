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

	fmt.Scan(&k)

	if calc(0, 0) {
		fmt.Println("ok")
	}
}

func calc(iter int, sum int) bool {

	if iter == n {
		return sum == k
	}

	if calc(iter+1, sum) {
		return true
	}

	if calc(iter+1, sum+a[iter]) {
		return true
	}

	return false
}
