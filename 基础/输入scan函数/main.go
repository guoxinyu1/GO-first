package main

import "fmt"

func main() {
	var name,age string
	fmt.Println("请输入您的姓名和年龄:")
	fmt.Scanln(&name,&age)
	//fmt.Scanf("%s\n%s",&name,&age)
	//fmt.Scanf("%s%s",&name,&age)
	fmt.Println("输入的内容是：",name,age)

}
