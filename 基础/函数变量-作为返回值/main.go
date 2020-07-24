package main

import "fmt"

func main() {
//调用a()-->返回值是一个函数
//-->result是一个函数变量
//函数变量不能直接输出,result要再次定义变量接受返回值
	result:=a()
	a:=result()
	fmt.Println(a)
	
}
//定义一个返回值是函数的无参函数
//func() int是函数a的返回值
func a() func() int  {
	return func() int {
		return 11
	}
}
