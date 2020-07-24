package main

import "fmt"

//参数不确定，等待接受main函数传递过来的参数。
func demo(i interface{}){
	//断言用法
	_,ok:=i.(int)
	if ok{
		fmt.Println("参数是int类型")
		return//return的是断言判断正确默认返回变量值，错误返回的是变量默认值
	}
	_,ok=i.(float64)
	if ok{
		fmt.Println("参数类型是float64")
		return
	}
	fmt.Println("参数既不是int也不是float64")


}

func main() {
	demo(false)
	
}
