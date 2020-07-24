package main

import "fmt"

func main() {

	//经典用法
	//for i:=0;i<5;i++{
	//	fmt.Println(i)
	//}


	//for循环遍历数组
	//arr:=[3]string{"guo","xin","yu"}
	//for i:=0;i<len(arr) ;i++  {
	//	fmt.Println(arr[i])
	//
	//}

	//for循环结合range函数
	//range arr-->会返回两个对象一个数组下标，一个数组元素
	//i是小标
	//n ：=arr[i]
	arr:=[3]string{"guo","xin","yu"}
	for i,n:=range arr{
		fmt.Println(i,n)

	}
	for _,n:=range arr{
		fmt.Println(n)

	}
}
