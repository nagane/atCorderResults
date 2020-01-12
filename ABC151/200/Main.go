package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var N, K, M int
	fmt.Scan(&N, &K, &M)
	A := make([]int, N-1)
	sc.Split(bufio.ScanWords)
	sum := 0
	for i := 0; i < N-1; i++ {
		sc.Scan()
		var e error
		A[i], e = strconv.Atoi(sc.Text())
		sum += A[i]
		if e != nil {
			fmt.Print(e)
		}
	}
	border := M * N
	if (border - sum) > K {
		fmt.Println("-1")
	} else if (border - sum) <= 0 {
		fmt.Println("0")
	} else {
		fmt.Println(border - sum)
	}
}
