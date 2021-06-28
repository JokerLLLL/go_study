package main

import (
	"fmt"
)

type workChan struct {
	in    chan int
	out chan bool
}

func works(i int, wc workChan) {
	for n:= range wc.in{
		fmt.Printf("work id：%d get channel info: %c int:%d \n", i, n, n) // %c 字符
		go func() {
			wc.out<-true
		}()
	}
}

func createWorks(i int) workChan {
	wc := workChan{
		in:    make(chan int),
		out: make(chan bool),
	}
	go works(i, wc)
	return wc
}

func chanDoneDemo() {
	//var c chan int // nil == c 无法使用
	var channels [10]workChan
	for i := 0; i < 10; i++ {
		channels[i] = createWorks(i)
	}

	for i := 0; i < 10; i++ {
		channels[i].in <- 'a' + i // 'a' 是ASCII码 == int
	}

	for i := 0; i < 10; i++ {
		channels[i].in <- 'A' + i
	}

	for _, wc := range channels {
		<-wc.out
		<-wc.out
	}

}

func main()  {
	chanDoneDemo()
}