package main

import "fmt"

// Mover 接口
type Mover interface {
	move()
}

type dog struct {
	name string
}

type car struct {
	brand string
}

// dog类型实现Mover接口
func (d dog) move() {
	fmt.Printf("%s会跑\n", d.name)
}

// car类型实现Mover接口
func (c car) move() {
	fmt.Printf("%s速度70迈\n", c.brand)
}

func main() {
	var x Mover            //接口类型
	a := dog{name: "旺财"}   //实例化结构体
	b := car{brand: "保时捷"} //实例化结构体
	x = a                  //a=结构体dog x=interface x内的接口类型来约束a结构体的成员函数的名和实现方法
	x.move()
	x = b
	x.move()
}
