package tryDefer

import (
	"bufio"
	"fmt"
	"github.com/jokerl/go_study/functional/fib"
	"os"
)

func WriterFile(filename string)  {
	file, e := os.Create(filename)
	//file, e := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	// e = errors.New("customer errors") //other error
	if e != nil {
		if error, ok := e.(*os.PathError); ok {
			fmt.Printf("%s,%s,%s", error.Op, error.Path, error.Err)
		} else {
			panic(e)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i:= 0 ; i < 100 ;i++ {
		defer fmt.Fprintln(writer, f()) // 逆序写入 //bug 100次 f() 越界了
	}
	fmt.Println("end writer")
}
