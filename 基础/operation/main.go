package main

import "fmt"

func main() {
	//算术运算符

	//a,b,c:=1,2,3
	//fmt.Println(a+b*c)
	//a++
	//fmt.Println(a)
	//a--
	//fmt.Println(a)
	////取余判断数字是否能被整除
	//fmt.Println(a%b)

	/*
	位运算:整数的原反补都相同不变
			负数的原码是负数的绝对值的二进制表示
			负数的反码是原码按位取反
			负数的补码是反码+1
	*/
	//原码：0000 0011
	//反码：0000 0011
	//补码：0000 0011
	var a int8=3
	//原码：0000 0010
	//反码：1111 1101
	//补码：1111 1110
	var b int8=-2
	fmt.Println(a,b)
	//正数和负数的位运算的左移和右移操作的都是源码,负数右移后不要忘了加负号。
	fmt.Println(a<<2)//12
	fmt.Println(b<<2)//-8
	fmt.Println(a>>2)//0
	fmt.Println(b>>1)//-1
	// |    &    ^（不同为1，相同为0.）   &^(位清空)-->正数和负数的操作的都是补码。
	fmt.Println(a|b)//1111 1111-->1111 1110-->0000 0001-->-1
	fmt.Println(a&b)//0000 0010-->2
	fmt.Println(a^b)//1111 1101-->1111 1100-->0000 0011-->-3
	fmt.Println(a&^b)//0000 0001-->1
}
