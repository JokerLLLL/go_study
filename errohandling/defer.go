package main

import (
	"fmt"
	"github.com/jokerl/go_study/errohandling/tryDefer"
)

func main() {
	tryDefer.DeferSomeThing()
	fmt.Println("file out put begin")
	tryDefer.WriterFile("defer1.txt")
	fmt.Println("file out put end")
}