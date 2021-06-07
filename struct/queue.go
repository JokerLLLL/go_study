package main

import (
	"fmt"
	"github.com/jokerl/go_study/struct/queue"
)

func main()  {
	q := queue.Queue{1,2,3}
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Push("6x6")
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q)
}