package main

import (
	"fmt"
	"sync"
	"time"
)

func countSelf(i int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		count := 0
		for {
			count++
			if (count >= int(time.Microsecond)) {
				fmt.Printf("协程：%d 完成：%d \n", i, count)
				break
			}
		}
		wg.Done()
	}()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		countSelf(i, &wg)
	}
	wg.Wait()
}
