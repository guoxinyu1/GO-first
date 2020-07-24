package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan string) //存放数据
	ch2 := make(chan int)    //同步
	go func() {
		//使用for循环将数据存入ch1
		for i := 97; i < 97+26; i++ {
			//ch1 <- strconv.Itoa(i)
			ch1<-fmt.Sprintf("%c",i)//将i转换成字符串并赋值。
		}
		ch2 <- 1//必须放在有效代码最后，不然阻塞后面的代码执行。

	}()
	go func() {
		//使用for range取出数据
		for n := range ch1 {
			fmt.Println(n)

		}

	}()
	<-ch2
	fmt.Println("程序结束。")
}
