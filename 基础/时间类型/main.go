package main

import (
	"fmt"
	"time"
)

func main() {
	//声明时间变量
	//var t time.Time
	//fmt.Println(t)

	//获取当前时间
	//t:=time.Now()
	//fmt.Println(t)

	//通过纳秒时间戳创建时间变量
	//t1:=time.Unix(0,t.UnixNano())
	//fmt.Println(t1)

	//自己指定时间
	// t2:=time.Date(2020,7,9,23,8,30,0,time.Local)
	// fmt.Println(t2)
	// t2.Month()

	//时间类型转换成字符串
	t:=time.Now()
	t1:=t.Format("2006-01-02 15:04:05")
	fmt.Println(t1)

}
