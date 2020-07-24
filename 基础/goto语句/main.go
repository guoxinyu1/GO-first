package main

import "fmt"

func main() {
	fmt.Println("执行程序")
	i:=6
	if i==6{
		goto Loop
	}
	fmt.Println("if下面输出")
Loop:
	fmt.Println("loop")


	//goto跳出循环
	for i:=0;i<5;i++{
		fmt.Println(i)
		if i==2 {
			goto  abc

		}
	}
	fmt.Println("循环结束")
abc:
	fmt.Println("abc")
	fmt.Println("程序结束")
}
