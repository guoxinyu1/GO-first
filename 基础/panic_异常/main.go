package main

import (
	"fmt"
)

func main() {

	//os.Exit(0)//立即终止程序。
	defer fmt.Println("遇到panic异常，我也要执行，我来收尾工作")
	fmt.Println("11111")
	panic("异常抛出，后面的代码不执行。")
	fmt.Println("222222")
}
