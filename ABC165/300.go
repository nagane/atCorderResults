package main

import "fmt"

func main() {
	var N, M, Q int
	fmt.Scan(&N, &M, &Q)
	d := make([][]int, Q)

	for i := 0; i < Q; i++ {
		d[i] = make([]int, 4)
		fmt.Scan(&d[i][0], &d[i][1], &d[i][2], &d[i][3])

	}
}
