package main

import (
	"fmt"
	"strings"
)

func main() {
	s:="smallming"
	//索引 Index(字符串,"字符")
	//	   LastIndex(字符串,"字符")
	fmt.Println(strings.Index(s,"l"))
	fmt.Println(strings.LastIndex(s,"l"))

	//判断是否指定内容开头或结尾返回结果是布尔
	fmt.Println(strings.HasPrefix(s,"small"))
	fmt.Println(strings.HasSuffix(s,"ming"))
	//判断内容是否包含指定字符串
	fmt.Println(strings.Contains(s,"mi"))
	//大小写
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	//替换
	fmt.Println(strings.Replace(s,"s","g",-1))
	//重复
	fmt.Println(strings.Repeat(s,10))
	//去掉空格或指定字符
	fmt.Println(strings.Trim(s," "))
	fmt.Println(strings.TrimSpace(s))
	//根据指定字符把字符串拆分成切片
	fmt.Println(strings.Split(s,"m"))
	fmt.Printf("%T\n",strings.Split(s,"m"))
	//join合并字符串切片，并且可以指定分割符
	arr :=[]string{"small","ming"}
	fmt.Println(strings.Join(arr,"*"))


}
