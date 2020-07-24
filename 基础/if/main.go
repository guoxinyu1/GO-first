package main

import "fmt"

func main() {
	score:=82

	//if的单独使用

	//if score>=65 {
	//	fmt.Println("及格")
	//
	//}
	//
	//if score1:=50;score1<60 {
	//	fmt.Println("不及格")
	//
	//}


	//if else
	//if  score>=60{
	//	fmt.Println("及格")
	//
	//}else {
	//	fmt.Println("不及格")
	//
	//}

	//if的多层嵌套
	//if score>=60 {
	//	if score>=80 && score<=100 {
	//		fmt.Println("优秀")
	//	}
	//	if score>=70 && score<80 {
	//		fmt.Println("中等")
	//	}
	//	if score>=60 && score<70 {
	//		fmt.Println("及格")
	//	}
	//
	//}else {
	//	fmt.Println("不及格")
	//}

	if score >=90{
		fmt.Println("优秀")
	}else if score>=80 {
		fmt.Println("中等")

	}else if score>=60 {
		fmt.Println("及格")

	}else {
		fmt.Println("不及格")
	}

}
