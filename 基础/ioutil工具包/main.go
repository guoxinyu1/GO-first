package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//file代表文件，来操作输入输出read(),write()一系列操作
	//file,err:=os.Open("D:/hello.txt")
	//defer file.Close()
	//if err!=nil{
	//	fmt.Println("文件打开失败",err)
	//	return//停止程序
	//}
	////读取内容,返回的是一个切片
	//b,_:=ioutil.ReadAll(file)
	//fmt.Println("文件中的内容:\n",string(b))

	//直接读取文件中的内容
	b, _ := ioutil.ReadFile("D:/hello.txt")
	fmt.Println(string(b))

	//写文件，会把原来的文件内容全部清空在写入。
	ioutil.WriteFile("D:/hello.txt", []byte("gxy"), 0666)
	fmt.Println("数据写入成功。")
	//获取文件中所有文件的信息
	f,_:=ioutil.ReadDir("D:/")
	//遍历切片f
	for _,n:=range f {
		fmt.Println(n.Name())

	}

}
