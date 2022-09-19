package main

import (
	"fmt"
	"time"
)

//结构体赋值
type Task struct {
	id         int
	name       string
	createTime time.Time //time.Time是结构体，属于结构体嵌套
	ok         bool
}

func main() {
	task := Task{id: 1, name: "sssss", ok: true, createTime: time.Now()}
	fmt.Println(task.ok)
}
