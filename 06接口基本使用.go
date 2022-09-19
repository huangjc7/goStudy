package main

import "fmt"

// Sayer 接口
type Sayer interface {
	say()
}

type dog struct {
}

//type cat struct {
//}

func (d *dog) say() string {
	return "汪汪汪"
}

//func (c *cat) say() string {
//	return "喵喵喵"
//}

func main() {
	fugui := dog{}
	dahuang := dog{}
	fmt.Println(fugui.say(), dahuang.say())
}
