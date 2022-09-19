package main

import (
	"fmt"
	"strconv"
)

func main() {
	log := [][4]string{
		{"1.1.1.1", "/index.html", "200", "1000"},
		{"1.1.1.2", "/index.html", "200", "10000"},
		{"1.1.1.1", "/index.html", "200", "10000"},
	}
	ip := map[string]int{}
	status := map[string]int{}
	ll := map[[2]string]int{} //map内嵌套数组 与key变量相互呼应

	for _, v := range log {
		//strconv.Atoi(v[0])
		ip[v[0]]++
		status[v[2]]++
		key := [2]string{v[0], v[1]}
		tr, _ := strconv.Atoi(v[3]) //重点 tr转为后，下面点+=不需要写tr[3]直接写tr即可
		ll[key] += tr
	}
	fmt.Println(ip)
	fmt.Println(status)
	fmt.Println(ll)

}
