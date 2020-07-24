package main

import "fmt"

func main() {
	result:=test()

	fmt.Println(result()) //2-->0xc00000a0b8
	fmt.Println(result()) //3-->0xc00000a0b8
	fmt.Println(result()) //4-->0xc00000a0b8
//重新调用时，重新声明变量，开辟新的内存地址。
	result2:=test()
	fmt.Println(result2()) //2-->0xc00000a108
	fmt.Println(result2()) //3-->0xc00000a108


	fmt.Println(result())//5-->0xc00000a0b8
//result 和 result2  指向同一个函数test（）内存地址，所以二者地址相同。
	fmt.Printf("%p %p \n",result,result2)


	
}
func test() func() int {
	i:=1
	return func() int {
		fmt.Printf("%p\n",&i)
		i+=1
		return i
	}

}
