package main

import "fmt"

func main() {
	//var a bool=true
	////var b bool=false
	//var c=true
	//d:=false
	//fmt.Println(a,c,d)
	//var e int8
	//var b byte=1
	//g:=2
	//fmt.Println(unsafe.Sizeof(e),unsafe.Sizeof(b),unsafe.Sizeof(b))
	//fmt.Println(unsafe.Sizeof(g))
	var a byte=1
	var b int8=0
	var c bool=true
	//bool不能和int和byte相互转换
	//a=byte(c)
	//b=int8(c)
	a=byte(b)
	fmt.Println(a,b,c)


}
