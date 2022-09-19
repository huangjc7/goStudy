package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var l sync.Mutex
var wg sync.WaitGroup

func add() {
	x++
	defer wg.Done()
}

func mutexAdd() {
	l.Lock()
	x++
	l.Unlock()
	defer wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&x, 1)
	defer wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go add() //5.328676ms
		//go mutexAdd() //耗时3.863921ms
		//go atomicAdd() //耗时2.872466ms
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
