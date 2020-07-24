package main

import "fmt"

func main() {
	//不接受返回值参数
	demo()
	//全部接受返回值
	a,b:=demo2()
	fmt.Println(a,b)
	//不希望接收的返回值使用下划线占位
	c,_:=demo2()
	fmt.Println(c)

	d,_:=demo3()
	fmt.Println(d)
	
}
func demo() (string,int) {
	return "smalling",17
}

func demo2()  (name string,age int){
	name="郭新宇"
	age=26
	return

}

func demo3() (string,int){
	return "xiaoming",18
}