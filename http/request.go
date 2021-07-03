package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"time"
)

func getImooc()  {
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bytes)
}

func main() {
	request, err := http.NewRequest("GET", "http://www.imooc.com", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Add("User-Agent","Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")

	cliect := http.Client{
		CheckRedirect:func(req *http.Request, via []*http.Request) error {
			// 处理重定向
			fmt.Println("redirect :", req.URL)
			return nil
		},
		Timeout:time.Second * 10,
	}
	resp, err := cliect.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", bytes)
}