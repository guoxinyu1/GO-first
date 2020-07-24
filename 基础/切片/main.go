package main

import "fmt"

func main() {
//切片就是指针，操作指针用*
	var slice []string
	fmt.Println(slice)
	fmt.Printf("%T",slice)

	//切片时引用类型
	names:=[]string{"g","x"}
	names1:=names
	names1[0]="zhang"
	fmt.Println(names1,names)
	fmt.Printf("%T",names1)

}
