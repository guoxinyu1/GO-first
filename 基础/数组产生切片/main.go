package main

import "fmt"

func main() {
	num:=[5]int{1,2,3,4,5}
	s:=num[0:2]
	fmt.Println(s)
	s[0]=0
	fmt.Println(s)
	fmt.Println(num)
	fmt.Printf("%p %p\n",s,&num)


	s=append(s,4,5,6,7)
	fmt.Println(s)
	fmt.Println(num)
	fmt.Printf("%p %p",s,&num)


//删除实现
num1:=[]int{0,1,2,3,4,5,6}
//要删除角标为n的元素
n:=2
num2:=num1[0:n]
num2=append(num2,num1[n+1:]...)
fmt.Println(num1)//0 1 3 4 5 6 6

}
