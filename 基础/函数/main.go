package main

import "fmt"


func main() {
	demo()
	a,b:=1,2
	add(a,b)
	sum1:=add2(a,b)
	fmt.Println(sum1)
	sum2:=add3(a,b)
	fmt.Println(sum2)

	
}
//无参函数
func demo()  {
	fmt.Println("执行demo")

}
//有参数函数
func add(c,d int)  {
	fmt.Println(c+d)

}

//有返回值的函数
func add2(c,d int) int  {
	return c+d

}

func add3(c,d int) (sum int) {
	sum=c+d
	return

}
