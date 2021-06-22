### go的unit test

1. 文件名必须以`_test.go`结尾 
2. 函数名称必须以`Test`开头 
3. 使用表格驱动测试，分离数据和测逻辑，并且错误不会终止程序 
4. `go test .` 子console上跑 
5. `go test -coverprofile=c.out` + `go tool cover -html c.out`
    命令行进行判断测试覆盖率


## go的bench test

1. 函数必须以`Benchmark`开头 
1. `go test -bench  .` 命令查看运行测试，for循环中使用b.N自动确认循环测试 

```textmate
goos: windows
goarch: amd64
pkg: github.com/jokerl/go_study/algo
cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
BenchmarkSearchLengthSonString-8         1228920               991.0 ns/op
PASS
ok      github.com/jokerl/go_study/algo 2.320s
```

2. 命令配置graphviz使用 

https://www.cnblogs.com/shuodehaoa/p/8667045.html   
https://graphviz.org/download  

命令：  
`go test -bench  . -cpuprofile cpu.out`  
`go tool pprof cpu.out`  