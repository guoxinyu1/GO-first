package main

import "fmt"

func main() {
	//go语言中，默认声明数据类型是int.
	e:=64

	var a int8=8
	var b int16=16
	var c int32=32
	var d int64=64
	fmt.Println(a,b,c,d)
	//类型转换 类型(值)
	fmt.Println(e+int(b))

	//go语言中不支持二进制赋值，但可以先转成字符串，在以二进制形式输出。
	b2 :=fmt.Sprintf("%b",a)
	fmt.Println(b2)

}
