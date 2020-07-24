package main

import (
	"fmt"
	"sync"
)

func main() {
	//声明读写锁变量
	var rw sync.RWMutex
	var wg sync.WaitGroup
	//m:=make(map[string]string)
	m2:=make(map[int]int)
	wg.Add(10)
	for i:=0;i<10 ;i++  {

		func(j int){
			//读写map里面的数据-->将整型参数通过调用strconv.Itoa()
			//赋值给map
			//m["key"+strconv.Itoa(j)]="value"+strconv.Itoa(j)
			//上锁
			rw.Lock()
			m2[j]=j
			fmt.Println(m2)
			//解锁
			rw.Unlock()
			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Println("程序结束")
	
}
