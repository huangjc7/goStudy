package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 1)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go test1(output1)
	go test2(output2)
	//使用select监控，读取到数据后执行
	//select可以同时监听一个或多个channel，直到其中一个channel ready
	select {
	case o1 := <-output1:
		fmt.Println("o1=", o1)
	case o2 := <-output2:
		fmt.Println("o2=", o2)
	}
}
