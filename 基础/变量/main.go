package main

import "fmt"

func main() {
	//1.先声明后赋值。
	var name string
	name="郭新宇"
	fmt.Println(name)
	//2.声明同时赋值,当两个变量赋值内容相同时，指向同一个内存地址。
	var name2 string="smallming"
	fmt.Println(name2)
	fmt.Printf("name2的内存地址是:%p\n",&name2)
	//3.声明同时赋值,省略类型
	var name3="smallming"
	fmt.Println(name3)
	fmt.Printf("name3的内存地址是:%p\n",&name3)
	//短变量只能用于函数体内。
	name4:="smallming"
	fmt.Println("短变量赋值方法:",name4)
	//多个变量声明方式：
	var(
		a=1
		b=2
		c=false
	)
	fmt.Println(a,b,c)
	//使用短变量声明并赋值多个变量时要求必须有一个变量没有声明
	var(
		s=1
		d=5
		f=3
	)
	m,d,f:=4,5,6
	fmt.Println(s,d,f,m)


}
