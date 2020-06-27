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
			maze[i][j] = str[j : j+1]
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
	for i := 0; i < N; i++ {
		d[i] = make([]int, M)
		for j := 0; j < M; j++ {
			d[i][j] = INF
		}
	}

	// スタート地点をキューに入れてその地点を0とする
	que := q.New(100000000)
	que.Put(P{x: sx, y: sy})
	d[sx][sy] = 0

	for que.Len() > 0 {
		p, _ := que.Get(1)
		var cur P
		cur = p[0].(P)

		// Getした座標がゴールなら抜ける
		if cur.x == gx && cur.y == gy {
			break
		}

		//移動4方向をループ
		for i := 0; i < 4; i++ {
			var nx, ny int //移動したあとの座標
			nx = cur.x + dx[i]
			ny = cur.y + dy[i]

			// 行ったことないならPut
			if 0 <= nx && nx < N && 0 <= ny && ny < M {
				if maze[nx][ny] != "#" && d[nx][ny] == INF {
					que.Put(P{nx, ny})
					// 移動できるならキューに入れ、その点の距離をpから距離+1で確定
					d[nx][ny] = d[cur.x][cur.y] + 1
				}
			}
		}
	}
	// ゴールまでのPOP
	print(d[gx][gy])
}
