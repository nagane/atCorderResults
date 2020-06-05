package main

import (
	"fmt"
	q "github.com/golang-collections/go-datastructures/queue"
)

func main() {
	que := q.New(10)
	fmt.Println(que)

}
