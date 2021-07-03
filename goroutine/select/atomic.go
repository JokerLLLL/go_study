package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type AtomicInt struct {
	value int
	m sync.Mutex // 通过 m实现线程安全
}

func (a *AtomicInt) increment() {
	a.m.Lock()
	defer a.m.Unlock()
	a.value ++
}

func (a *AtomicInt) print() {
	a.m.Lock()
	defer a.m.Unlock()
	fmt.Println(a.value)
}

func main()  {
	var a AtomicInt
	var a2 int64
	atomic.AddInt64(&a2, 2) // go 提供的atomic
	a.increment()

	go func() {
		a.increment()
		atomic.AddInt64(&a2, 2)
	}()
	time.Sleep(time.Millisecond)
	a.print()
}
