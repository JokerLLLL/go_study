package  main

import (
	"fmt"
	"time"
)

func work(i int, c <-chan int) {
	for n:= range c{
		fmt.Printf("work id：%d get channel info: %c int:%d \n", i, n, n) // %c 字符
	}
}

func createWork(i int) chan<- int {
	c := make(chan int) // mark 出来的channel <-c<-
	go work(i, c)
	return c
}

func chanDemo() {
	//var c chan int // nil == c 无法使用
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i // 'a' 是ASCII码 == int
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

}

func chanBufferDemo() {
	c := make(chan int, 3) // 3个buffer channel 没有消费者会报错死锁
	c<-'@'
	c<-98
	c<-99
	go work(0, c)
	c<-100
}

func chanCloseDemo() {
	c := make(chan int, 3) // 3个buffer channel 没有消费者会报错死锁
	c<-'@'
	c<-98
	c<-99
	go work(0, c)
	c<-100
	close(c) // 关闭
}

func main()  {
	chanDemo()
	//chanBufferDemo()
	//chanCloseDemo()
	time.Sleep(time.Second)
}