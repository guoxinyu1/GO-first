package main

import (
	"fmt"
	"strings"
)

func main() {
	//1个中文占三个字节，所以字符串长度12个字节。
	s:="smallming张"
	//fmt.Println(len(s))
	//获取字符串中的单个字符,字符(go不支持字符数据)都是int32的数字,先
	//sprintf()转换成字符串.
	//s1:=s[0]
	//fmt.Println(s1)
	//fmt.Printf("%c\n",s1)//此时是以字符格式化输出,s1还是uint类型.
	//fmt.Printf("%T\n",s1)
	//
	//s2:=fmt.Sprintf("%c\n",s1)
	//fmt.Println(s2)
	//fmt.Printf("%T",s2)


	//截取字符序列=字符串
	//s1:=s[9:12]
	//fmt.Println(s1)



	//切片获取
	//将占3个字节的中文字符,格式化长度为10
	//[]rune(变量)这样用
	//s2 := []rune(s)将字符串s转换成切片,并赋值给s2
	s2 := []rune(s)
	fmt.Println(len(s2))
	fmt.Printf("%c\n",s2[9])
	fmt.Printf("%T\n",s2[9])

	//遍历字符串
	for _,n:=range s{
		//fmt.Println(i)
		fmt.Printf("%c",n)

	}
	strings.ToLower(s)

}
