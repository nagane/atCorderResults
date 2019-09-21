package main

import (
	"os"
	"strconv"
)

func main() {

	h, _ := strconv.Atoi(os.Args[1])
	if h < 0 {
		os.Exit(0)
	}
	w, _ := strconv.Atoi(os.Args[2])
	if w > 1000 {
		os.Exit(0)
	}

	matrix := make([][]int, h)

	for i := 0; i < h; i++ {
		matrix[i] = make([]int, w)
		for j := 0; j < w; j++ {
			matrix[i][j] = 0
		}
	}

	for i := 0; i < h; i++ {
		print("\n")
		for j := 0; j < w; j++ {
			print(matrix[i][j])
		}
	}
}
