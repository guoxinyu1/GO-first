package main

import "fmt"

//定义结构体
type People struct{
	Name string
	Age int
}
func main() {

	//声明结构体变量
	var student People
	//赋值
	student=People{"张三",18}
	fmt.Println(student)

	teacher:=People{Name:"gxy",Age:18}
	fmt.Println(teacher)

	var mom People
	mom.Name="hh"
	mom.Age=25
	fmt.Println(mom)
	//判断 ==判断的是两个结构体变量中的内容是否相等
	student=People{"张三",18}
	student1:=People{"张三",18}
	fmt.Printf("%p %p\n",&student,&student1)
	fmt.Println(student==student1)

	
}
