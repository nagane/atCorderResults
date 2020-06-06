package main

import (
	"fmt"
	q "github.com/golang-collections/go-datastructures/queue"
)

func main() {
	que := q.New(10)
	fmt.Println(que)

	que.Put(1)
	que.Put(5)

	fmt.Println(que.Get(5))

}
