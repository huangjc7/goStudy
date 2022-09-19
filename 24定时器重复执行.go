package main

import (
	"fmt"
	"time"
)

func main() {
	// 1.获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	// 子协程
	i := 0
	go func() {
		for {
			fmt.Println(<-ticker.C)
			i++
			if i == 5 {
				ticker.Stop()
			}
		}
	}()
	for { //防止主历程关闭
	}
}
