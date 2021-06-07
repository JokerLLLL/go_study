package retriever

import "fmt"

type Retriever interface {
	Get(url string) string
	Post(url string) string
}

func Download(r Retriever) string  {
	return r.Get("http://www.baidu.com")
}

func Inspect(r Retriever)  {
	// name + struct
	fmt.Printf("%T %v \n", r, r)
	// type switch
	switch r.(type) {
	case MockRetriever:
		fmt.Println("this mock retriever")
	case *HttpRetriever:
		fmt.Println("this http retriever pointer")
	default:
		fmt.Println("other Retriever")
	}
}