package main

import (
	"fmt"
)

func main() {
	var N string
	fmt.Scan(&N)

	var point int

	switch len(N) {
	case 1:
		point = 0
	case 2:
		point = 1
	case 3:
		point = 2
	}

	switch string(N[point]) {
	case "0":
		fmt.Println("pon")
	case "1":
		fmt.Println("pon")
	case "2":
		fmt.Println("hon")
	case "3":
		fmt.Println("bon")
	case "4":
		fmt.Println("hon")
	case "5":
		fmt.Println("hon")
	case "6":
		fmt.Println("pon")
	case "7":
		fmt.Println("hon")
	case "8":
		fmt.Println("pon")
	case "9":
		fmt.Println("hon")
	}

}
