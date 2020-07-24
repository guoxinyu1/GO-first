package main

import "fmt"

func main() {
	//指针没有指向，没有分配地址空间。
	//var point *int
	//fmt.Printf("%p\n",point)

	//make函数初始化切片
	s:=make([]int,0,3)
	fmt.Printf("%p\n",s)

	//len（）cap（）
	fmt.Println(len(s),cap(s))

	//切片长度容量翻倍
	c:=make([]string,0)
	fmt.Println(len(c),cap(c))

	c=append(c,"gxy","学go")
	fmt.Println(len(c),cap(c))

	c=append(c,"很扭力")
	fmt.Println(len(c),cap(c))
	//如果添加多个值，且添加的长度大于翻倍后的容量，则容量和长度相等了。
	c=append(c,"4","5","6","7","8","9")
	fmt.Println(len(c),cap(c))

	//下次添加时，没有超过扩容后的容量，则只翻倍。
	c=append(c,"10")
	fmt.Println(len(c),cap(c))

	fmt.Println(c)

	//直接添加另一个切片
	d:=make([]int,0)
	f:=[]int{1,2,3}
	d=append(d,f...)
	fmt.Println(d)





}
