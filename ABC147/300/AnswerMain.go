package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	var A []int
	var x [][]int
	var y [][]int
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
		for j := 0; j < A[i]; j++ {
			fmt.Scan(&x[i][j], &y[i][j])
		}
	}

	ans := 0

	for bits := 1; bits < 1<<uint(N); bits++ {
		ok := true
		for i := 0; i < N; i++ {
			if 1 > (bits & (1 << uint(i-1))) {
				fmt.Println(ok)
				ans++
			}
		}

	}
}
