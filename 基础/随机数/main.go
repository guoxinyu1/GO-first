package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//使用当前时间的纳秒差保证种子每次运行时是不同的。
	rand.Seed(time.Now().UnixNano())
	//rand.Intn(num)生成一个10以内的0-9随机数
	fmt.Println(rand.Intn(10))

}
