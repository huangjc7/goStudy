package main

import "fmt"

func main() {
	var ch1 chan int   // 声明一个传递整型的通道
	var ch2 chan bool  // 声明一个传递布尔型的通道
	var ch3 chan []int // 声明一个传递int切片的通道
	// make(chan 元素类型, [缓冲大小])
	ch4 := make(chan int)
	ch5 := make(chan bool)
	ch6 := make(chan []int)
	fmt.Println(ch1, ch2, ch3, ch4, ch5, ch6)
}
