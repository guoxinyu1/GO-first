package main

import (
	"container/ring"
	"fmt"
)

func main() {

	//r代表整个循环链表，又代表第一个元素
	r:=ring.New(5)//创建了一个5个元素长度的双重循环链表
	r.Value=0
	r.Next().Value=1
	r.Next().Next().Value=2
	r.Prev().Value=4
	r.Prev().Prev().Value=3

	//循环链表有几个元素，匿名函数func就执行几次
	//i代表当前执行元素的内容
	r.Do(func(i interface{}) {
		fmt.Println(i)

	})//0 1 2 3 4
fmt.Println("*************move****************")
	//移动n次，支持负数
	fmt.Println(r.Move(0).Value)//0
	fmt.Println(r.Move(1).Value)//1
	fmt.Println(r.Move(-1).Value)//4

fmt.Println("*************添加Link()****************")
	r.Do(func(i interface{}) {
		fmt.Print(i," ")

	})
fmt.Println()
	r1:=ring.New(1)
	r1.Value=5
	//r.Link(r1)//添加到当前元素（第一个元素）的后面
	r.Next().Link(r1)
	r.Do(func(i interface{}) {
		fmt.Print(i," ")

	})
fmt.Println("*************删除Ulink()****************")
r.Do(func(i interface{}) {
		fmt.Print(i," ")

	})
fmt.Println()
//删除当前元素后面的一个元素删除
//当前元素r是不能被删的.
	//r.Unlink(1)
//把当前元素后面的两个元素一块删除
	//r.Unlink(2)

	//Unlink(取值)-->n%r.len()
	//即r.Unlink(2)和r.Unlink(8)效果一样
	r.Unlink(8)


r.Do(func(i interface{}) {
		fmt.Print(i," ")

	})
}

