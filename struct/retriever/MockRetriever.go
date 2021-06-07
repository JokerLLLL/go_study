package retriever

type MockRetriever struct {
	Contents string
}

func (r MockRetriever) String() string {
	return r.Contents
}

func (r MockRetriever) Get(url string) string {
	return r.Contents
}
func (r MockRetriever) Post(url string) string  {
	return r.Contents
}

