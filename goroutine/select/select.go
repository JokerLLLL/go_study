package main

import (
	"fmt"
	"math/rand"
	"time"
)


func work(i int, c chan<- int) {
	var j int
	for {
		time.Sleep(time.Duration(rand.Intn(1500 * i)) * time.Millisecond)
		c <- j
		j++
	}
}

func doWork() chan int {
	c := make(chan int)
	go func() {
		for i:= range c {
			time.Sleep(time.Duration(1 * time.Second)) // 模拟处理接收数据过慢
			fmt.Printf("get info: %d \n", i)
		}
	}()
	return c
}

func createWork(i int) chan int {
	c := make(chan int)
	go work(i, c)
	return c
}
//  csp模型
func main() {
	var c1, c2 = createWork(1), createWork(2)
	var do = doWork()
	var queue []int
	var timeAfter = time.After(20 * time.Second)
	var timeTick = time.Tick(5 * time.Second)

	for {
		var activeChan chan int // nil channel 永远杜塞住
		var v int
		if (len(queue) > 0) {
			activeChan = do
			v = queue[0]
		}
		select {
		case n := <-c1:
			queue = append(queue, n)
		case n := <-c2:
			queue = append(queue, n)
		case activeChan <- v:
			queue = queue[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Printf("get number time out! \n")
		case <-timeTick:
			fmt.Println(queue)
			time.Sleep(time.Second) //只影响当前case里的代码 说明也是goroutine？
		case <-timeAfter:
			fmt.Println(queue)
			fmt.Printf("goodbye")
			return
		}
	}
}
