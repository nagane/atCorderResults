package main

import "fmt"

func main() {
	var C string
	fmt.Scan(&C)
	nextChar := int(byte(C[0])) + 1
	fmt.Println(string(nextChar))
}
