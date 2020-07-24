package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	//客户端向服务器端发送数据
	//1.创建服务器端地址
	addr,_:=net.ResolveTCPAddr("tcp4","localhost:8899")

	for i:=0;i<5 ;i++  {
		//2.创建连接对象--本地地址和服务器地址进行连接。
		conn,_:=net.DialTCP("tcp",nil,addr)
		//3.连接对象写入数据并发送
		conn.Write([]byte("你好！我是来自客户端的消息gxy"+strconv.Itoa(i)))//三次握手连接

		/*4.接受服务器传递回来的数据
		(1).创建字节切片准备接受服务器返回的数据
		(2).conn.read()读取数据
		*/
		b:=make([]byte,256)
		c,_:=conn.Read(b)
		fmt.Println("接收到服务器的消息：",string(b[:c]))

		//5.关闭连接
		conn.Close()//四次握手关闭

	}

	fmt.Println("客户端结束")

	
}


