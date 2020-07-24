package main

import "fmt"

func main() {
	a:=new(int)
	fmt.Println(a)//输出指针地址
	*a=123
	fmt.Println(*a,a)



}
