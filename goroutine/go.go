package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ { // i 最后是 11
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
				//fmt.Printf("goroutine No. %d \n", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a) // a 和 a[i]++ data race 冲突
}
