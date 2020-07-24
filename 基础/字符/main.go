package main

import "fmt"

func main() {
	//定义数字
	var i rune=24352
	fmt.Println(i)
	//输出汉字张
	fmt.Printf("%c\n",i)
	//获取转换后的内容
	c:=fmt.Sprintf("%c\n",i)
	fmt.Println(c)
	g:=0x90ED
	//fmt.Println(g)
	fmt.Printf("%c\n",g)
}
