package main

import "fmt"

func main() {
	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	//遍历读取管道数据，如果data是具体数据，ok是布尔值，如果有数据返回true
	for {
		if data, ok := <-ch; ok { //ok是布尔值，如果有数据返回true
			fmt.Println(data)
		} else {
			break
		}
	}
}
