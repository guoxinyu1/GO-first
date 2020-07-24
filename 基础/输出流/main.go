package main

import (
	"fmt"
	"os"
)

func main() {
	filePath:="D:/hello.txt"
	//已可读写的方式打开文件
	file,err:=os.OpenFile(filePath,os.O_APPEND,0666)
	defer file.Close()
	if err!=nil{
		file,err=os.Create(filePath)
	}
	b:=[]byte("hello word\r\n")
	file.Write(b)
	file.WriteString("hello\nword")
	fmt.Println("输出流运行成功。")
}
