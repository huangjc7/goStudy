package main

import "fmt"

func main() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"

	for i := 1; i <= 5; i++ {
		fmt.Println(map1[i])
	}
}
