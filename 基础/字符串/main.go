package main

import (
	"fmt"
	"strconv"
)

func main() {
	//s:="small\nming"
	//s2:=`small\nming`
	//fmt.Println(s,s2)
	//s:="3000"
	//i,_:=strconv.ParseInt(s,10,64)
	//i,_:=strconv.Atoi(s)
	//fmt.Println(i)
	//fmt.Printf("%T",i)

	//把int转换成字符串
	//var i1 int =17
	//s1:=strconv.FormatInt(int64(i1),10)
	//fmt.Println(s1)
	//fmt.Printf("%T",s1)
	//s:=strconv.Itoa(i1)
	//fmt.Println(s)
	//fmt.Printf("%T",s)

	//string转换成float
	//s:="10.5"
	//f,_:=strconv.ParseFloat(s,64)
	//fmt.Println(f+1)
	//fmt.Printf("%T",f)


	//将浮点型转换成字符型
	f:=1.522200000000000000
	s:=strconv.FormatFloat(f,'f',6,64)
	fmt.Println(s)



}
