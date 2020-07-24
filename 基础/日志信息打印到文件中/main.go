package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file,err:=os.OpenFile("D:/golog.log",os.O_APPEND|os.O_CREATE,0777)
	if err!=nil{
		fmt.Println("日志文件打开失败：",err)
		return
	}
	logger:=log.New(file,"时间：",log.Ltime)
	logger.Println("输出日志信息")
	fmt.Println("程序结束")
	
}
