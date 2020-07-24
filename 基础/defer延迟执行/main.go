package main

import "fmt"

//defer与return
func f()  (i int){
	i=0
	defer func() {
		i+=2
	}()
	return
}

func main() {
	//单个defer使用
	//fmt.Println("打开链接")
	//defer func() {
	//	fmt.Println("关闭连接")
	//
	//}()
	//fmt.Println("进行操作")

	//多个defer-》先产生的后持行
	//fmt.Println("打开链接a")
	//defer func() {
	//	fmt.Println("关闭连接a")
	//
	//}()
	//fmt.Println("进行操作b")
	//defer func() {
	//	fmt.Println("6666")
	//}()
	//fmt.Println("fff")

	fmt.Println(f())

	
}
