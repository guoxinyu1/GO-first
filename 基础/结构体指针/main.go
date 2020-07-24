package main

import "fmt"

type People struct {
	Name string
	Age int
}
func main() {
	peo:=new(People)//指针，且有地址
	peo.Name="佳明哥"//peo是指针，但其结构体中属性不是指针

	peo1:=peo//peo1指针，二者地址相同
	peo1.Name="555"
	fmt.Println(peo1,peo)

//两个指针p1,p2,即使内容相同但是地址两块。
	p1:=&People{"S",1}
	p2:=&People{"S",1}
	fmt.Println(p1==p2)//false
	fmt.Println(*p1==*p2)//true
}
