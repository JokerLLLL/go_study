package fileService

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const prefix  = "/list/"


// 用户错误接口 和 用户错误实现
type UserError interface {
	error
	Message() string
}

type RouterUserError string

func (e RouterUserError) Error() string  {
	return string(e)
}

func (e RouterUserError) Message() string {
	return string(e.Error())
}


type AppHandel func(writer http.ResponseWriter, request *http.Request) error

func AppHandelReal(writer http.ResponseWriter, request *http.Request) error {

	// panic("test")

	if strings.Index(request.URL.Path, prefix) != 0 {
		return RouterUserError("router muster begin "+ prefix)
	}

	path := request.URL.Path[len(prefix):]
	file, e := os.Open(path)
	if e != nil {
		return e
	}
	defer file.Close()
	bytes, e := io.ReadAll(file)
	if e != nil {
		return e
	}
	writer.Write(bytes)
	return e
}

func ErrorHandel(handel AppHandel) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// catch all panic
		defer func() {
			if err := recover(); err != nil {
				// TODO err to string 怎么转换
				log.Printf("Panic: %v", err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		e := handel(writer, request)
		if e != nil {
			httpCode := http.StatusOK
			// handel user error
			if ue, ok := e.(UserError); ok {
				http.Error(writer, ue.Message(), http.StatusBadRequest)
				return
			}
			// handel sys error
			switch {
			case os.IsNotExist(e):
				httpCode = http.StatusNotFound
			case os.IsPermission(e):
				httpCode = http.StatusForbidden
			default:
				httpCode = http.StatusInternalServerError
			}
			log.Printf("Err: %s", e.Error())
			http.Error(writer, http.StatusText(httpCode), httpCode)
		}
	}
}