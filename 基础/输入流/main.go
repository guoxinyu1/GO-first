package main

import (
	"fmt"
	"os"
)

func main() {
	//使用string包下的NewReader创建字符串流
	//返回的是reader指针类型
	//reader:=strings.NewReader("hello 世界")
	////reader转换成字节切片。数据流读取read([] byte)
	//b:=make([] byte,reader.Size())
	////把流中数据读取到切片中。
	//n,err:=reader.Read(b)
	//if err!=nil{
	//	fmt.Print("读取失败",err)
	//	return
	//}
	//fmt.Println("读取数据长度为：",n)
	//fmt.Println("流中数据",string(b))


	//文件流
	f,err:=os.Open("D:/hello/hello.txt")//打开文件
	if err!=nil{
		fmt.Println("文件读取失败",err)
		return
	}
	fileInfo,err:=f.Stat()//获取文件信息
	if err!=nil{
		fmt.Println("文件信息读取失败",err)
		return
	}
	//根据文件中数据大小创建字节切片，准备放入到read()
	b:=make([] byte,fileInfo.Size())
	//把文件f中的数据读取到切片b中。
	n,err:=f.Read(b)
	if err!=nil{
		fmt.Println("文件流读取失败",err)
		return
	}
	fmt.Println("文件流中数据长度为：",n)
	fmt.Println("文件内容为：",string(b))


}
