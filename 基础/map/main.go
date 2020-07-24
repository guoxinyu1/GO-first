package main

import "fmt"

func main() {
	//map完整写法，此时没有分配内存地址
	var a map[int]int
	fmt.Println(a)
	fmt.Printf("%p\n",a)
	//make函数实例化map，分配内存空间

	m:=make(map[int]int)
	fmt.Println(m==nil)
	fmt.Printf("%p\n",m)

	//一行和多行的语法区别
	m1:=map[int]int{1:001,2:002,3:003}
	m2:=map[int]int{
		1:001,
		2:002,
		3:003,
	}
	fmt.Println(m1,m2)

	//key不存在新增数据，存在覆盖数据
	salary:=map[string]int{"小明":1000,"gxy":2000}
	salary["gxy"]=68000
	fmt.Println(salary)
	salary["小弟"]=68952314
	salary["小k"]=6
	fmt.Println(salary)
	//delete()顶层函数调用删除元素
	delete(salary,"gxy")
	fmt.Println(salary)
   //map中要删除的key不存在，map中的内容不变也不会报错。
	delete(salary,"gxy3")
	fmt.Println(salary)

	//查询，通过key进行操作
	//key存在返回对应的value值
	//不存在返沪value类型的默认值
	fmt.Println(salary["小明"])
	fmt.Println(salary["gxy6"])

	//value:key对应的值，ok表示key是否存在。
	_,ok:=salary["小明"]
	fmt.Println(ok)
	value,ok:=salary["小明"]
	fmt.Println(value,ok)

	//循环遍历map所有内容
	for key,value:=range salary {
		fmt.Println(key,value)

	}





}
