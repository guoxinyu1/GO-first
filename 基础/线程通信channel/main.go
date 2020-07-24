package main

import "fmt"

/*
同步：
channel实现main主协程和子协程go func之间通信。
*/
//func main() {
//	ch:=make(chan int)
//	go func() {
//		fmt.Println("进入goroutine")
//		ch<-1//必须在最后
//	}()
//	a:=<-ch
//	fmt.Println(a)
//	fmt.Println("程序结束")
//
//}
func main()  {
	//声明一个通道来阻塞主线程，给goroutine运行时间。
	ch2:=make(chan int)
	ch1:=make(chan string)
	go func() {
		fmt.Println("进入goroutine1")
		ch1<-"hello我是goroutine1"
		ch2<-1//同步必须放在有效代码最后一行
}()
	go func() {
		fmt.Println("进入goroutine2")
		//声明一个变量接受通道里的数据
		content:=<-ch1
		fmt.Println("取出数据：",content)
		ch2<-2//同步必须放在有效代码最后一行

	}()
	//阻塞主协程
	<-ch2
	<-ch2

	fmt.Println("程序结束")
}
