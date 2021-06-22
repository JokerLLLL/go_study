package queue

import "fmt"

// go 语言的实例代码
func ExampleQueue() {
	q := Queue{1, 2, 3}
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	q.Push("6x6")
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q)
	//Output:
	//1
	//2
	//false
	//3
	//true
	//6x6
	//<nil>
	//[]
}