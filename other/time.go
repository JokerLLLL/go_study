package main

import (
	"fmt"
	"strconv"
	"time"
)

func main()  {
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Nanosecond())
	fmt.Println(strconv.FormatInt(time.Now().UnixNano(), 10))
}
