package tryDefer

import (
	"fmt"
	"os"
)

// tryDefer 在函数调用时候运行 在行首
func DeferSomeThing()  {
	defer fmt.Println("last ")
	defer fmt.Println("first")
	fmt.Println("this is function first")
	for i:=0;i<5;i++ {
		defer fmt.Printf("i ：%d \n", i)
	}
	os.Create("sss")
	fmt.Printf("xxx")

}
