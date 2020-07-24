package main

import "fmt"

func main() {
	//死锁
	//ch:=make(chan int)
	//ch<-1
	//fmt.Println("死锁")


	//有缓存的channel
	ch:=make(chan int,3)
	ch<-1
	<-ch
	ch<-1
	ch<-1
	ch<-1

	fmt.Println("程序结束")
	
}
