package main

import "fmt"

func main() {
	//a:=3
	//fmt.Printf("变量a的内存地址是%p,代表开辟了一块内存空间。\n",&a)
	//fmt.Println("变量a也就是内存中存储的数据是:",a)
	//a=4
	//fmt.Printf("a被重新赋值修改的是内容，地址没变：%p\n",&a)
	//fmt.Println("a内容发生变化：",a)
	//
	//b:=a
	//b=5
	//fmt.Println("b开辟一块新的内存空间：",&b,&a)
	//创建两个变量，整型和指针变量。内存空间有两块。
	a:=123
	var point *int
	point=&a
	fmt.Println(point)
	//.Println(&point)


	*point=3
	fmt.Println(*point,a)
	fmt.Println(point)
	//fmt.Println(&point)





}
