package main

import (
	"bufio"
	"fmt"
	"github.com/jokerl/go_study/functional/fib"
)

func main() {
	f := fib.Fibonacci()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	// interface了Read接口
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
