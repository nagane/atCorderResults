package main

import (
	"fmt"
	q "github.com/golang-collections/go-datastructures/queue"
)

func main() {
	que := q.New(10)

	que.Put(1)
	fmt.Println(que)
	que.Put(5)

	fmt.Println(que)
	fmt.Println(que.Get(5))
	fmt.Println(que)

}
