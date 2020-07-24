package main

import "fmt"

func main() {
	//a,b是实参
	a:=1
	b:="郭新宇"
	demo(a,b)//生成值的副本传入demo（）方法中
	fmt.Println(a,b)//不变
	
}
//i=1， s="郭新宇"-->新开辟出两个内存空间存贮副本
func demo(i int,s string)  {
	i=5
	s="xiaoming"
}
