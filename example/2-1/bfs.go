package main

import (
	"fmt"
)

//BFS: Breadth-First Search
func main() {
	var N, M int
	fmt.Scan(&N, &M)
	maze := make([][]string, N)

	//迷路インプット
	for i := 0; i < N; i++ {
		var str string
		fmt.Scan(&str)
		maze[i] = make([]string, M)
		for j := 0; j < M; j++ {
			maze[i][j] = str[:j]
		}
	}

}
