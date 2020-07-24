package main

import "fmt"

func main() {
	//for i:=0;i<5 ;i++  {
	//	if i==2||i==4{
	//		continue
	//	}
	//	fmt.Println(i)
	//
	//}

	//双层for循环默认影响的是最内层循环
	//for i:=0;i<2 ;i++  {
	//	for j:=0;j<3 ;j++  {
	//		if i==1{
	//			continue
	//		}
	//		fmt.Println(j,i,"循环结束")
	//
	//	}
	//
	//}

//自定义循环标签，让continue语句控制该标签。
	myfor:for i:=0;i<2 ;i++  {
		for j:=0;j<3 ;j++  {
			if i==1{
				continue myfor
			}

			fmt.Println(j,i,"循环结束")

		}

	}
	fmt.Println("完全结束")
}
