package main

import (
	"fmt"
	"sync"
	"time"
)

//读写锁 增加判断，如果是读操作不会加锁，如果是写操作会加锁
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	//lock.Lock() // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()              // 解写锁
	//lock.Unlock() // 解互斥锁
	wg.Done()
}

func read() {
	//lock.Lock() // 加互斥锁
	rwlock.RLock() // 加读锁
	//time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	time.Sleep(time.Millisecond)
	rwlock.RUnlock() // 解读锁
	//lock.Unlock() // 解互斥锁
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
