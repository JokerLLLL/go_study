package main

import (
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func helloWorld(writer http.ResponseWriter, request *http.Request)  {
	io.WriteString(writer, "Hello World")
}

func main() {
	http.HandleFunc("/d", helloWorld)
	e := http.ListenAndServe(":8989", nil)
	if e == nil {
		log.Fatalln(e)
	}
}
