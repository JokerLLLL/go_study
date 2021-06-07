package main

import (
	"fmt"
	"github.com/jokerl/go_study/functional/adder"
)

func main()  {
	a := adder.Adder() // 计数器
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	a2 := adder.Adder2(0) // 累加器
	value := 0
	for i := 0; i < 10; i++ {
		value, a2 = a2(i)
		fmt.Printf("0+1+...+%d = %d \n", i, value)
	}

}
