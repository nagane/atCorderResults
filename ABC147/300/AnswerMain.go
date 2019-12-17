package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	x := make([][]int, N)
	y := make([][]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
		x[i] = make([]int, A[i])
		y[i] = make([]int, A[i])
		for j := 0; j < A[i]; j++ {
			fmt.Scan(&x[i][j], &y[i][j])
		}
	}

	ans := 0

	for bits := 1; bits < (1 << uint(N)); bits++ {
		ok := true
		for i := 0; i < N; i++ {
			if bits&(1<<uint(i)) == 0 {
				continue
			}
			for j := 0; j < A[i]; j++ {
				fmt.Println(((bits >> uint(x[i][j]-1)) & 1) ^ y[i][j])
				if ((bits>>uint(x[i][j]-1))&1)^y[i][j] == 1 {
					ok = false
				}
			}
		}
		if ok {
			ans = max(ans, counter(bits))
		}
	}
	fmt.Println(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func counter(x int) int {
	if x == 0 {
		return 0
	}
	return counter(x>>1) + (x & 1)
}
