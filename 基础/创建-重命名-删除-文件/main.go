package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		创建文件夹：1.os.Mkdir()要求文件夹不存在且父目录必须存在。
		2.os.Mkdir()如果文件夹存在也不报错，保留原文件夹。
					如果父目录不存在就自动创建


	*/
	//D:/-->父目录必须存在
	//	err:=os.Mkdir("D:/godir",os.ModeDir)
	//	if err!=nil{
	//		fmt.Println("文件夹创建失败，错误信息：",err)
	//		return
	//	}
	//	fmt.Println("文件夹创建成功。")
	//如果文件夹存在也不报错，保留原文件夹。
	//	err := os.MkdirAll("D:/godir/hello", os.ModeDir)
	//	if err != nil {
	//		fmt.Println("文件夹创建失败，错误信息：", err)
	//		return
	//	}
	//	fmt.Println("文件夹创建成功。")

	/*
	   创建空文件:1.要求文件目录必须存在
	   			2.如果文件存在则会创建一个空文件覆盖之前的文件。
	   			3.os.Create()
	*/
//	file, err := os.Create("D:/godir/hello/hello.txt")
//	if err != nil {
//		fmt.Println("文件创建失败，错误信息：", err)
//
//	} else {
//		fmt.Println(file.Name(), "文件创建成功。")
//}


		/*
		重命名文件和文件夹：os.Rename("","")
		1.第一个参数，原文件夹或文件的名称，带路径
		2.第二个参数，新的名称，带路径。


		*/
	//err := os.Rename("D:/godir/hello/hello.txt","D:/godir/hello/hello2.txt")
	//if err != nil {
	//	fmt.Println("文件重命名失败，错误信息：", err)
	//
	//} else {
	//	fmt.Println( "文件重命名成功。")
	//}

	//获取文件（夹）的信息
	//1.f,err:=os.open()打开文件
	//2.msg,err:=f.stat()获取文件信息
	//3.msg.Name()-->获取属性
	//4.文件打开后还要关闭 defer f.close()
	//file,err:=os.Open("D:/godir/hello/hello2.txt")
	//defer file.Close()
	//if err!=nil {
	//	fmt.Println("打开文件失败，错误信息为：",err)
	//	fmt.Println(&err)
	//	return
	//}
	//fileInfo,err:=file.Stat()
	//if err!=nil {
	//	fmt.Println("获取文件信息失败，错误信息为：",err)
	//	fmt.Println(&err)
	//	return
	//}
	//fmt.Println(fileInfo.Name())
	//fmt.Println(fileInfo.IsDir())//是否是文件夹，true，false
	//fmt.Println(fileInfo.Mode())//文件权限
	//fmt.Println(fileInfo.ModTime())//修改时间
	//fmt.Println(fileInfo.Size())//文件大小


	//删除
	err:=os.RemoveAll("D:/godir")
	if err!=nil{
		fmt.Println("文件删除失败",err)
		return
	}
	fmt.Println("删除成功")
}
