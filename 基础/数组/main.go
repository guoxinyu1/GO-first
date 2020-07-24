package main

import "fmt"

func main() {
	//声明数组的方式。
	//var arr1 [3]int=[3]int{1,2,3}
	//arr2:=[3]int{1,2,3}
	//arr3:=[10]int{1}
	//fmt.Println(arr1,arr2,arr3)
	//arr4:=[...]int{1}
	//fmt.Println(arr4,len(arr4))

	//数组是值类型。
	arr1:=[3]int{1,2,3}
	arr2:=arr1
	fmt.Println(arr1,arr2)
	fmt.Printf("%p %p\n",&arr1,&arr2)
	arr2[0]=2
	fmt.Println(arr1,arr2)
	fmt.Printf("%p %p",&arr1,&arr2)


}
