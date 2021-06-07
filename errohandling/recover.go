package main

import (
	"fmt"
)

func tryRecover() {
	//panic("string error")
	//panic(errors.New("this is an error"))
	b := 0
	fmt.Println(5 / b)
}

func catchErrorAndHandel()  {
	i := recover() // 必须在defer中调用，否则是nil
	if err,ok := i.(error) ; ok {
		fmt.Println(err.Error())
	}else{
		panic(i)
	}
}

func main() {
	defer catchErrorAndHandel()
	tryRecover()
}
