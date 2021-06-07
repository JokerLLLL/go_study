package main

import (
	"fmt"
	"github.com/jokerl/go_study/struct/retriever"
	"time"
)

func main()  {


	// 想用 type assertion 就必须var定义变量 否则报错
	var r retriever.Retriever
	r = retriever.MockRetriever{"this mock retriever"}
	retriever.Inspect(r)

	// type assertion
	if Mr,ok := r.(retriever.MockRetriever); ok {
		fmt.Println(Mr.Contents)
	}

	var r2 retriever.Retriever
	r2 = &retriever.HttpRetriever{"Mozilla/5.0", time.Minute}

	//println(retriever.Download(r))
	//println(retriever.Download(r2))

	retriever.Inspect(r)
	retriever.Inspect(r2)

	r3 := retriever.HttpRetriever{}
	retriever.Inspect(&r3)
}