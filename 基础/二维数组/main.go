package main

import "fmt"

func main() {
	//var arr [3][3]int=[3][3]int{{1,2,3},{1,2,3},{1,2,3}}
	//var arr2 [2][2]int=[2][2]int{{0,1},{0,1}}
	//fmt.Println(arr2,arr)

//二维数组。
	arr:= [3][3]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	fmt.Println(arr)

	fmt.Println(arr[0])//{1,2,3}
	arr0:=arr[0]
	fmt.Println("我是一个数组",arr0)
	fmt.Println(arr[0][0],arr[0][1],arr[0][2])


	//三位数组
	arr3:=[2][2][2]int{
		{
			{1,2},
			{3,4},
		},
		{
				{5,6},
				{7,8},
		},
	}
	fmt.Println(arr3[1][0][0])
	
}
