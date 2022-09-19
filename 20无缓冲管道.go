package main

import "fmt"

func recv(c chan int) {
	x := <-c
	fmt.Println("接收成功", x)
}

func main() {
	ch := make(chan int)

	go recv(ch)

	ch <- 10
	fmt.Println("发送成功")
}
