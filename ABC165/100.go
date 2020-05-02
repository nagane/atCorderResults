package main

import "fmt"

func main() {
	var K, A, B int
	fmt.Scan(&K)
	fmt.Scan(&A, &B)

	i := 1
	for ; i <= A; i++ {
		if (i * K) >= A {
			break
		}
	}
	if (i * K) <= B {
		fmt.Println("OK")
	} else {
		fmt.Println("NG")
	}
}
