package main

import "fmt"

func main() {
	//默认表示十进制
	d:=17
	//o开头表示八进制
	o:=021
	//0x开头表示十六进制
	x:=0x11
	//en表示10的n次方
	e:=11e2
	//输出
	fmt.Println(d,o,x,e)
	//go语言中不支持二进制赋值，但可以先转成字符串，在以二进制形式输出。
	b :=fmt.Sprintf("%b",d)
	fmt.Println(b)
}
