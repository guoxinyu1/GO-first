package main

import "fmt"

func main() {
	//当某个条件是固定值
	//switch num:=1;num {
	//default:
	//	fmt.Println("内容不正确。")
	//case 2:
	//	fmt.Println("二进制")
	//case 8:
	//	fmt.Println("八进制")
	//case 10:
	//	fmt.Println("10进制")
	//case 16:
	//	fmt.Println("16进制")
	//}
//当条件是范围时
//score:=60
//	switch  {
//	default:
//		fmt.Println("不及格。")
//	case score>=90:
//		fmt.Println("优秀")
//	case score>=80:
//		fmt.Println("中等")
//	case score>=70:
//		fmt.Println("良")
//	case score>=60:
//		fmt.Println("及格")
//	}


//case后面可以跟多个条件
//month:=2
//	switch month {
//	case 1,3,5,7,8,10,12:
//		fmt.Println("31天")
//	case 2,11:
//		fmt.Println("28/29天")
//	default:
//		fmt.Println("default")
//	}


//穿透
//	switch num:=6;num {
//	default:
//		fmt.Println("default")
//		fallthrough
//	case 1:
//		fmt.Println("1")
//
//	case 2:
//		fmt.Println("2")
//	case 3:
//		fmt.Println("3")
//
//	}
	//break
	switch num:=9;num {
	default:
		fmt.Println("default")
		fallthrough
	case 1:
		break
		fmt.Println("1")

	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")

	}

	fmt.Println("程序结束。")
	
	
}
