package main

import "fmt"

func main() {

	//匿名函数无参数无返回值
	func(){
		fmt.Println("无参数无返回值的匿名函数")

	}()
//有参数的匿名函数
	func(name string){
		fmt.Println("名字是:",name)
	}("郭新宇")

	//有返回值的匿名函数
	name:=func() (name string){
		name="更新"
		return
	}()//此处有调用
	fmt.Println(name)
}
