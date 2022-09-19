package main

func main() {
	//1.timer基本使用
	//timer1 := time.NewTimer(2 * time.Second) //设置一个定时器 2秒触发
	//t1 := time.Now()                         //获取当前时间
	//fmt.Printf("t1:%v\n", t1)                //打印当前时间
	//t2 := <-timer1.C                         //时间到则触发
	//fmt.Printf("t2:%v\n", t2)

	// 2.验证timer只能响应1次
	//timer2 := time.NewTimer(time.Second)
	//for {
	//	<-timer2.C
	//	fmt.Println("时间到")
	//}

	//小实验：重复定时器触发

	//time := time2.NewTimer(time2.Second * 1)
	//
	//for {
	//	t1 := <-time.C
	//	fmt.Println(t1)
	//}

	// 3.timer实现延时的功能
	//(1)
	//time.Sleep(time.Second)
	//(2)
	//timer3 := time.NewTimer(2 * time.Second)
	//<-timer3.C
	//fmt.Println("2秒到")
	//(3)
	//<-time.After(2*time.Second)
	//fmt.Println("2秒到")

	// 4.停止定时器
	//timer4 := time.NewTimer(2 * time.Second)
	//go func() {
	// <-timer4.C
	// fmt.Println("定时器执行了")
	//}()
	//b := timer4.Stop()
	//if b {
	// fmt.Println("timer4已经关闭")
	//}

	// 5.重置定时器
	//timer5 := time.NewTimer(3 * time.Second)
	//timer5.Reset(1 * time.Second)
	//fmt.Println(time.Now())
	//fmt.Println(<-timer5.C)

	for {

	}
}
