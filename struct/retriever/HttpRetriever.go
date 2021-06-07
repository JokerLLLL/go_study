package retriever

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type HttpRetriever struct {
	UserAgent string
	TimeOut time.Duration
}

func (r HttpRetriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	b, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	_ = resp.Body.Close()
	return string(b)

}

// retriever.HttpRetriever does not implement retriever.Retriever (Post method has pointer receiver)
// 只要有一个是 pointer 就只能 type assertion *HttpRetriever
func (r *HttpRetriever) Post(url string) string {
	return url
}