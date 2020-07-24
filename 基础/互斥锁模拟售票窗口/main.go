package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	//票数
	num = 100
	//等待组完成协程并发功能
	wg sync.WaitGroup
	//互斥锁保证数据num安全，保证只有一个协程访问它。
	mu sync.Mutex
)
//声明售票函数
func sellTicker(i int)  {
	defer wg.Done()
	for{
		//加锁，保证num被当前的协程读取。
		mu.Lock()
		if num>=1{
			fmt.Printf("第%d个窗口卖了%d\n",i,num)
			num-=1
		}
		//解锁
		mu.Unlock()
		if num<0{
			break
		}
		//wg.Done()
		time.Sleep(time.Duration(rand.Int63n(1000)*1e6))
	}

}
func main() {
	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	for i:=1;i<=4;i++{
		wg.Add(1)
		go sellTicker(i)
	}
	wg.Wait()
	fmt.Println("程序结束,所有票卖完")

}
