package main

import (
	"fmt"
	q "github.com/golang-collections/go-datastructures/queue"
)

// P is interface
type P struct {
	x int
	y int
}

//BFS: Breadth-First Search
func main() {
	var N, M int
	fmt.Scan(&N, &M)
	maze := make([][]string, N) //迷路を表す文字列の配列

	//スタート、ゴールの座標
	var sx, sy, gx, gy int

	//迷路インプット
	for i := 0; i < N; i++ {
		var str string
		fmt.Scan(&str)
		maze[i] = make([]string, M)
		for j := 0; j < M; j++ {
			// TODO: インプットついでにスタートとゴールの座標を取りたい
			maze[i][j] = str[:j]
			//　スタート位置を記憶
			if maze[i][j] == "S" {
				sx = i
				sy = j
			}
			if maze[i][j] == "G" {
				gx = i
				gy = j
			}
		}
	}

	// INF
	INF := 100000000
	//移動ベクトル
	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}

	d := make([][]int, N) //各点までの最短距離の配列
	for i := 0; 1 < N; i++ {
		d[i] = make([]int, M)
		for j := 0; j < M; j++ {
			d[i][j] = INF
		}
	}

	// スタート地点をキューに入れてその地点を0とする
	que := q.New(100000000)
	que.Put(P{sx, sy})

	for que.Len() > 0 {

	}

}
