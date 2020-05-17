package main

import (
	"fmt"
)

func main() {
	var K int
	var S string
	fmt.Scan(&K)
	fmt.Scan(&S)

	if len(S) > K {
		str := S[:K] + "..."
		fmt.Println(str)

	} else {
		fmt.Println(S)
	}

}
