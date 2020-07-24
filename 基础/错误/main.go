package main

import (
	"errors"
	"fmt"
)

//使用go语言标准库创建错误，并返回。
//e=errors.New()
func demo(i, k int) (d int, e error) {

	if k == 0 {
		//使用errors包下下的New()返回一条错误信息
		e = errors.New("除数不能为0")
		d = 0
		return

	}
	d = i / k
	return
}

func demo1(i, k int) (d int, e error) {
	if k == 0 {
		//返回的错误信息由多个内容组成
		//e =fmt.Errorf("除数不能是0，两个参数分别是：%d %d",i,k)
		e =fmt.Errorf("%s %d 和 %d","除数不能是0，两个参数分别是：",i,k)
		d = 0
		return

	}
	d = i / k
	return

}

func main() {
	result, error := demo1(6, 0)
	//fmt.Println(result, error)
	if error!=nil{
		fmt.Println("程序执行出错，错误信息",error)
		return
	}
	fmt.Println("程序执行成功，结果为：",result)

}
