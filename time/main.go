package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	now := time.Now()
	s := now.Unix()      //秒
	ns := now.UnixNano() //纳秒

	fmt.Println(now)      //当前时间格式
	fmt.Println(s)        //
	fmt.Println(ns)       //纳秒
	fmt.Println(ns / 1e6) //纳秒转毫秒
	fmt.Println(ns / 1e9) //纳秒转秒
	fmt.Println()
	fmt.Println(time.Unix(s, 0).String())                            //秒转时间格式
	fmt.Println(time.Unix(0, ns).Format("2006-01-02 15:04:05.999999999 -0700 MST"))                           //纳秒时间格式
	fmt.Println(time.Unix(s, 0).Format("2006-01-02 15:04")) // 秒转时间格式 并时间格式化

	// groupId: 1 start: 2021-10-07 16:09:38.4953881 +0800 CST  end: 2021-10-07 16:09:38.4953881 +0800 CST duration: 1000 ms
	var logs []string
	for i := 0; i < 5; i++ {
		group := "0_" + strconv.FormatInt(int64(i), 10)
		s := time.Now()
		time.Sleep(time.Second)
		e := time.Now()
		logs = append(logs, Timelog(group, s, e))
	}

	PrintLog(logs)

}

func Timelog(groupId string, start time.Time, end time.Time) string {
	s := time.Unix(0, start.UnixNano()).Format("2006-01-02#15:04:05.999999999")
	e := time.Unix(0, end.UnixNano()).Format("2006-01-02#15:04:05.999999999")
	d := (end.UnixNano() - start.UnixNano()) / 1e6

	return fmt.Sprintf("groupId: %s start: %s  end: %s duration: %d ms", groupId, s, e, d)
}

func PrintLog(s []string)  {
	for _,i:= range s {
		println(i)
	}
}