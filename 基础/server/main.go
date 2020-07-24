package main

import (
	"fmt"
	"net"
)

func main() {
	//1.创建自己的服务器的地址
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	//2.创建监听器，监视端口
	lis, _ := net.ListenTCP("tcp4", addr)
	for {
		//3.通过监听器获取客户端传递过来的数据
		//阻塞式,一直等待数据传入，然后才能进行下一步运行
		conn, _ := lis.Accept()
		//添加子协程gorutinue提高运行效率
		go func() {
			//4.读取数据-->读取到一个字节切片中。
			b := make([]byte, 1024)
			n, _ := conn.Read(b)
			fmt.Println("接受到客户端的数据为:", string(b[:n]))
			/*向客户端返回数据--响应*/
			//5.连接对象写入数据
			conn.Write([]byte("服务器:我在."))
			//6.关闭连接
			conn.Close()
		}()
	}
	fmt.Println("服务器结束")

}
