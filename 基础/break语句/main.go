package main

import "fmt"

func main() {
	//for i:=0;i<3 ;i++  {
	//	for j:=0;j<3 ;j++  {
	//		if j==1{
	//			break
	//		}
	//		fmt.Println(i,j)
	//
	//	}
	//
	//}
	myfor:for i:=0;i<3 ;i++  {
		for j:=0;j<3 ;j++  {
			if j==1{
				break myfor
			}
			fmt.Println(i,j)

		}

	}


	fmt.Println("程序结束")

	
}
