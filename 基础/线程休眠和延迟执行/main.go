package main

import (
	"fmt"
	"time"
)

func main() {
	//线程休眠
	//fmt.Println("程序开始")
	//time.Sleep(10e9)//程序阻塞10秒
	//fmt.Println("程序结束")

	//延迟执行
	fmt.Println("程序开始")
	//2秒后执行匿名函数
	time.AfterFunc(2e9, func() {
		fmt.Println("延迟触发")
	})
	time.Sleep(4e9)//程序阻塞4秒，等待延迟触发
	fmt.Println("程序结束")

	
}
