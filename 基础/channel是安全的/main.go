package main

import (
	"fmt"
	"time"
)

func main() {
	ch:=make(chan int)
	for i:=1;i<5;i++  {
		go func(j int) {
			fmt.Println(j,"开始")

			fmt.Println(j,"结束")
			ch<-j
		}(i)

	}
	for j:=1;j<5 ;j++  {
		time.Sleep(2*time.Second)
		<-ch
	}
	fmt.Println("程序结束")
	
}
