## go package && go mod



1. go package始终加载不了

http://c.biancheng.net/view/5394.html

GO111MODULE 作怪打扰了GOPATH的加载
https://blog.csdn.net/qq_38151401/article/details/105729884
https://www.cnblogs.com/pu369/p/12068645.html

2. mod

http://c.biancheng.net/view/5712.html  

https://www.jianshu.com/p/760c97ff644c  

```
go env

go get github.com/sirupsen/logrus
```

windows的`GOPATH`可直接在环境变量设置

##  检查goroutine数据访问冲突

go run -race go.go 检测数据访问冲突

$ go run -race go.go
==================
WARNING: DATA RACE
Read at 0x00c00012e058 by goroutine 7:
  main.main.func1()
      D:/demo/go/src/go_study/goroutine/go.go:14 +0x77

Previous write at 0x00c00012e058 by main goroutine:
  main.main()
      D:/demo/go/src/go_study/goroutine/go.go:11 +0x144

Goroutine 7 (running) created at:
  main.main()
      D:/demo/go/src/go_study/goroutine/go.go:12 +0x115
==================
==================
WARNING: DATA RACE
Read at 0x00c000108070 by goroutine 7:
  main.main.func1()
      D:/demo/go/src/go_study/goroutine/go.go:14 +0x9c

Previous write at 0x00c000108070 by goroutine 8:
  main.main.func1()
      D:/demo/go/src/go_study/goroutine/go.go:14 +0x4e

6.goroutine切换
  1 .I/O select
  2. channel
  3. 锁
  4. 函数调用(有时）
  5. runtime.Gosched()

### go build main.go 需要其他.go文件

https://blog.csdn.net/coolboyzero/article/details/77946653?utm_source=blogxgwz0

###  go mod 

如何使用
https://www.jianshu.com/p/bbed916d16ea

1.获取当前包的依赖，2.设置编辑器支持mod
https://blog.csdn.net/weixin_30436891/article/details/95664229


####  http服务

在httpsevice起来后，只调用HttpClient() 一次初始化
后续所有的http链接（）

pacakge client

func HttpClient() *http.Client {
	fmt.Println("request client!!!!")
	if httpClient == nil {
		httpClient = &http.Client{}
		httpClient.Timeout = 30 * time.Second
		//连接池设置
		httpClient.Transport = &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second, // 连接超时时间
				KeepAlive: 120 * time.Second, // 保持长连接的时间
			}).DialContext,                          // 设置连接的参数
			MaxIdleConns:          100,              // 最大空闲连接
			IdleConnTimeout:       60 * time.Second, // 空闲连接的超时时间
			ExpectContinueTimeout: 30 * time.Second, // 等待服务第一个响应的超时时间
			MaxIdleConnsPerHost:   20,              // 每个host保持的空闲连接数
		}
		fmt.Println("init http client with pool!!!")
	}
	return httpClient
}

## Writer和Reader的含义

type Writer interface {
   Write(p []byte) (n int, err error)
}

  从`p`里读取数据(数据源），向底层(本身这个实例)写入数据。

type Reader interface {
   Read(p []byte) (n int, err error)
}

 读取底层内容(其实就是本身这个实例）来源不管，向`p`里写入内容。 然后给到你使用。
 
 
 ## 其他
 
  prinln
 
 println() 和 fmt.Println()  不再同个线程 输出有先后（待确认）
 
  goland快捷键
 
 https://www.cnblogs.com/just-save/p/12993705.html
 
 
  协程
 
 非抢占式多任务处理：协程 由协程主动交出控制权
 抢占式任务处理：线程
 
 
 ## go-http-json的用法
 
 https://blog.csdn.net/weixin_36750623/article/details/119936014