package main

import (
	"runtime"
	"testing"
	"time"
)

//go test *.go -v可看到数据
func TestAll(t *testing.T) {
	go StartAPI()//如需测试data race请按照README.md要求切换到master分支然后注释此行代码
	testCases := []string{
		"1,2",
		"1,2",
		"clear",
		"1,2,3",
		"1,2,4",
	}
	runtime.GOMAXPROCS(1)
	for _, ts := range testCases {
		if ts == "clear" {
			ClientHttp("", "/clzclear")
		} else {
			go ClientHttp(ts, "/clzarr")
			if ts == "1,2" { //如果等于1,2就多执行一次并发
				go ClientHttp(ts, "/clzarr")
			}
		}
	}
	time.Sleep(time.Duration(5) * time.Second)
}
