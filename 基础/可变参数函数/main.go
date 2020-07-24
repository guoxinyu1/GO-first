package main

import "fmt"

func main() {
	a,b,c:="1","23","20"
	demo(a,b,c)

	
}

func demo(hover...string){

	//i是下标
	//n是切片内容.
	hover[0]="sb"
	for i,n:=range hover  {
		fmt.Println(i,n)

	}



}

