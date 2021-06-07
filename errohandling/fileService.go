package main

import (
	"github.com/jokerl/go_study/errohandling/fileService"
	"net/http"
)

func main() {
	http.HandleFunc("/", fileService.ErrorHandel(fileService.AppHandelReal))
	e := http.ListenAndServe(":8888", nil)
	if e != nil {
		panic(e)
	}
}
