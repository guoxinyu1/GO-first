package main

import "fmt"

func demo1()  {
	fmt.Println("demo1上半部分")
	demo2()
	fmt.Println("demo1下半部分")

}
func demo2()  {
	defer func() {
		recover()
	}()
	fmt.Println("demo2上半部分")
	demo3()
	fmt.Println("demo2下半部分")

}
func demo3()  {

	fmt.Println("demo3上半部分")
	panic("demo3中出现panic")
	fmt.Println("demo3下半部分")

}
func main() {
	fmt.Println("程序开始")
	demo1()
	fmt.Println("程序结束")
}
