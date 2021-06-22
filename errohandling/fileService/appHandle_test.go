package fileService

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func AppHandelPanic(writer http.ResponseWriter, request *http.Request) error {
	panic("this panic handle")
}
func AppHandleNotFound(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrNotExist
}
func AppHandleForbidden(writer http.ResponseWriter, request *http.Request) error {
	return os.ErrPermission
}
func AppHandelUnknown(writer http.ResponseWriter, request *http.Request) error {
	return errors.New("unknown handle err")
}

func AppHandelUserError(writer http.ResponseWriter, request *http.Request) error {
	return RouterUserError("router muster begin list")
}
func AppHandelSuccess(writer http.ResponseWriter, request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	writer.WriteHeader(http.StatusOK)
	return nil
}

var tests = []struct {
	handle  AppHandel
	code    int
	message string
}{
	{AppHandelPanic, 500, "Internal Server Error"}, // this panic handle 错误隐藏
	{AppHandleNotFound, 404, "Not Found"},
	{AppHandleForbidden, 403, "Forbidden"},
	{AppHandelUnknown, 500, "Internal Server Error"},
	{AppHandelUserError,400, "router muster begin list"},
	{AppHandelSuccess,200, "no error"},
}

func TestErrorHandel(t *testing.T) {

	for _, test := range tests {
		var writer = httptest.NewRecorder() // 必须在for里面 因为一个writer只能写一次Code！！！
		var request = httptest.NewRequest(http.MethodGet, "http://test.com", nil)
		f := ErrorHandel(test.handle)
		f(writer, request)

		resp := writer.Result() // or writer.Body == writer.Result().Body
		checkResponseCase(resp, test, t)
	}
}

func TestAppHandelRealByServer(t *testing.T) {
	for _, test := range tests {
		var server = httptest.NewServer(http.HandlerFunc(ErrorHandel(test.handle)))
		resp, _ := http.Get(server.URL)
		checkResponseCase(resp, test, t)
	}
}

func checkResponseCase(resp *http.Response, testCase struct{handle AppHandel;code int;message string}, t *testing.T)  {
	b, _ := ioutil.ReadAll(resp.Body)

	msg := strings.Trim(string(b), "\n")
	if resp.StatusCode != testCase.code || msg != testCase.message {
		// 优化：1·test.handle 利用反射 获取出真正的handle名
		t.Errorf("handel %v, expected (%d,%s), actual (%d,%s)", testCase.handle, testCase.code, testCase.message, resp.StatusCode, msg)
	}
}
