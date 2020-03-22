package main

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	A := make([]int, N+M)
	lenCount := 0

	nCount := 0
	mCount := 0

	for i := 1; ; i++ {
		if i%2 == 0 && nCount < N {
			A[lenCount] = i
			nCount += 1
			lenCount += 1
		}
		if i%2 != 0 && mCount < M {
			A[lenCount] = i
			mCount += 1
			lenCount += 1
		}
		if lenCount > len(A)-1 {
			break
		}
	}

	anser := 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			if (A[i]+A[j])%2 == 0 {
				anser += 1
			}
		}
	}
	fmt.Println(anser)
}
