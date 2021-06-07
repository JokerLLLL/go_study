package fib

import (
	"fmt"
	"io"
	"strings"
)

type fibo func() int

func (f fibo) Read(p []byte) (n int, err error) {
	next := f()
	if next > 20000 {
		return  0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// TODO byte 和 strings.NewReader(s)  不能byte去读 fix
	return strings.NewReader(s).Read(p)

}

func Fibonacci() fibo {
	value1, value2 := 0, 1
	return func() int {
		value1, value2 = value2, value1+value2
		return value1
	}
}