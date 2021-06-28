package main

import (
	"fmt"
	"sync"
)

type workChan struct {
	in   chan int
	out  chan bool
	wg   *sync.WaitGroup
	done func()
}

func works(i int, wc workChan) {
	for n:= range wc.in{
		fmt.Printf("work id：%d get channel info: %c int:%d \n", i, n, n) // %c 字符
		wc.done()
	}
}

func createWorks(i int, wg *sync.WaitGroup) workChan {
	wc := workChan{
		in: make(chan int),
		wg: wg,
		done: func() {
			wg.Done()
		},
	}
	go works(i, wc)
	return wc
}

func chanDoneDemo() {
	//var c chan int // nil == c 无法使用
	var channels [10]workChan
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		channels[i] = createWorks(i, &wg)
	}

	wg.Add(20)
	for i, wc := range channels {
		wc.in <- 'a' + i // 'a' 是ASCII码 == int
	}
	for i, wc := range channels {
		wc.in <- 'A' + i
	}
	wg.Wait()
}

func main()  {
	chanDoneDemo()
}