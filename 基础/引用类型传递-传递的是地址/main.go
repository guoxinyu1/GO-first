package main

import "fmt"

func main() {
	a:=[]int{0,1,2}
	demo(a)
	fmt.Println(a)//[4 5 6]

	p:=7
	demo2(&p)
	fmt.Println(p)//8
	
}
func demo(slice []int)  {
	slice[0]=4
	slice[1]=5
	slice[2]=6

}

func demo2(point *int)  {
	*point=8

}

