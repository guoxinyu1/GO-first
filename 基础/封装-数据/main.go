package main

import (
	"fmt"
	"go_moudle/封装-数据/结构体"
)

//跨包访问结构体内的属性-->封装数据
//通过方法来访问。

func main() {
	//这样是获取不到结构体内首字母小写的属性
	//student:=结构体.People{"gxy",15}
	//fmt.Println(student)

	var student 结构体.People//结构体是包名
	student.SetName("gxy")
	student.SetAge(20)
	name:=student.GetName()
	fmt.Println(name)
	age:=student.GetAge()
	if age==0{
		fmt.Println("年龄设置不正确。")
	}else {
		fmt.Println("年龄为:",age)
	}

	
}
