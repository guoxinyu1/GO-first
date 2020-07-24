package main

import "fmt"

func main() {

	//函数作为参数
	//调用mydo时对参数进行声明.
	mydo(func(name string) {
		fmt.Println(name)
	})

}

//定义了一个函数mydo
//其中它的参数是一个有参数的函数变量-->arg func(name string)
//

func mydo(arg func(name string)) {
	fmt.Println("执行mydo")
	//调用mydo的参数-->函数变量arg()
	//此处调用arg func(name string),并没有进行声明
	//所以要在调用mydo()-->对arg func(name string)进行声明.
	arg("你好")


	//函数作为返回值

}
