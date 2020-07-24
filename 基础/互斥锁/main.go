package main

import (
	"fmt"
	"sync"
)

var (
	num = 100
	wg  sync.WaitGroup
	l  sync.Mutex
)
func demo(){
	l.Lock()//上锁，其他协程访问不了
	for i:=0;i<10;i++{
		num-=1
	}
	l.Unlock()//解锁
	wg.Done()
}
func main() {
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go demo()
	}
	wg.Wait()
	fmt.Println(num)

}
