package main

import (
	"fmt"
	"sort"
)

func main() {
	//排序
	name := []int{2, 324, 2, 12, 1}
	sort.Ints(name)
	fmt.Println(name)

	names := []string{"yz", "ac", "hj"}
	sort.Strings(names)
	fmt.Println(names)

	//go提供的二分查找方法
	name = []int{1, 3, 10, 20, 100}
	fmt.Println(name[sort.SearchInts(name, 10)] == 10)

	name = []int{1, 3, 10, 20, 100, 101, 150, 160, 170, 180, 187, 200, 201, 390}
	target := 10
	left := 0
	right := len(name)
	for left <= right {
		mid := left + (right-left)/2
		if target == name[mid] {
			fmt.Println(mid)
			break
		}
		if target > name[mid] {
			left = mid + 1
			continue
		}
		if target < name[mid] {
			right = mid - 1
			continue
		}
	}

	a := []int{123, 232, 12, 64, 7, 200}
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
			fmt.Println(i, a)
		}
	}
	fmt.Println(a)

}
