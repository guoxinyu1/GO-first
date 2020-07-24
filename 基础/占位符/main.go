package main

import (
	"fmt"
)

func main() {
	//输出结果,不会换行。
	fmt.Printf("%d %x %X %o %b",5,10,10,10,10)
	//换行
	fmt.Println()
	//将转换后的结果获取到
	a:=fmt.Sprintf("%x",18)
	fmt.Println(a)
	fmt.Printf("%f",1.5)
	fmt.Println()
	//内存地址
	i:=6
	fmt.Printf("%p",&i)

}
