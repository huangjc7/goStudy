package main

import "fmt"

type Task struct {
	name    string
	shengao int
}

type shuaiqi interface {
	start()
}

func newTask(name string, shengao int) Task {
	t := Task{name, shengao}
	return t
}

func (t *Task) start() {
	fmt.Printf("帅气逼人的%v,身高%d", t.name, t.shengao)
}

func main() {
	var x shuaiqi             //实例化接口 定义start()方法
	dx := newTask("黄健晨", 188) //通过newTask函数 return 参数进入Task结构体
	x = &dx                   // 实例化的Task结构体赋予接口，实现接口约束
	x.start()                 //启动函数
}
