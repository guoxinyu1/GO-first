package main

import "fmt"

func main() {
	//存数据
	ch := make(chan int)
	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}
	//select结合for循环取数据
	for i := 0; i < 10; i++{
		select {
		case c := <-ch:
			fmt.Println(c)
		//default:
		}

	}
}
