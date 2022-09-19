package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var x int
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() //加锁
		x = x + 1
		lock.Unlock() //解锁
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
