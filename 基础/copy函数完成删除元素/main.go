package main

import "fmt"

func main() {
	//使用cope函数完成删除元素，保证原切片内容不变
	s:=[]int{1,2,3,4,5,6}
	n:=1
	newslice:=make([]int,n)
	copy(newslice,s[0:n])
	newslice=append(newslice,s[n+1:]...)
	fmt.Println(s)
	fmt.Println(newslice)//完成删除

	
}
