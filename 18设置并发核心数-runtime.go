package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2) //两个任务只有一个逻辑核心，此时是做完一个任务再做另一个任务。
	// 将逻辑核心数设为2，此时两个任务并行执行
	go a()
	go b()
	time.Sleep(time.Second)
}
