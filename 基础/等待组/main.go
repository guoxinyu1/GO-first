package main

import (
	"fmt"
	"sync"
)

func demo(index int)  {
	for i:=1;i<=5 ;i++  {
		fmt.Printf("第%d执行的值为：%d\n",index,i)

	}
	wg.Done()

}
//创建等待组变量
var wg sync.WaitGroup
func main() {
	for i:=1;i<=5 ;i++  {
		wg.Add(1)
		go demo(i)
	}
	wg.Wait()//time.Sleep(3e9)
	fmt.Println("程序结束")
	
}
