package main

import (
	"fmt"
	"sort"
)

func main() {
	//对int类型切片排序
	num:=[]int{1,2,384,6}
	sort.Ints(num)//升序，默认升序
	fmt.Println(num)
	//降序  需要将切片转换成data interface数据类型
	sort.Sort(sort.Reverse(sort.IntSlice(num)))
	fmt.Println(num)


	//对float64类型切片排序

	f:=[]float64{1.5,7.2,5.8,2.3,6.9}
	sort.Float64s(f)
	fmt.Println(f)
	//降序
	sort.Sort(sort.Reverse(sort.Float64Slice(f)))
	fmt.Println(f)


	//对string类型切片排序
	s:=[]string{"a","郭新宇","gxy","dasd"}
	//升序
	sort.Strings(s)
	fmt.Println(s)
   //查找目标字符并返回下标。如果没有该字符则返回应插入的下标对应位置。
	n:=sort.SearchStrings(s,"郭新宇")
	fmt.Println(n)
	//降序
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)


	//逻辑
	sExample :=[]string{"还","you","谁"}
	//查找目标元素，切片必须首先进行升序
	sort.Strings(sExample)
	fmt.Println(sExample)
	numberIndex :=sort.SearchStrings(sExample,"you")
	fmt.Println(numberIndex)
	if numberIndex <len(sExample){
		if sExample[numberIndex]=="you"{
			fmt.Println("存在")
		}else{
			fmt.Println("不存在")
		}

	}else{
		fmt.Println("不存在")
	}
}
