package main

import "fmt"

func justifyType(x interface{}) {
	//断言类型
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func main() {
	//空接口
	justifyType(1111)

	// 定义一个空接口x  接口是一种类型
	var x interface{}
	s := "pprof.cn"
	x = s //无论是空接口或者定义接口都需要将结构体或者具体的值赋予接口，方可生效
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i //无论是空接口或者定义接口都需要将结构体或者具体的值赋予接口，方可生效
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b //无论是空接口或者定义接口都需要将结构体或者具体的值赋予接口，方可生效
	fmt.Printf("type:%T value:%v\n", x, x)

}
