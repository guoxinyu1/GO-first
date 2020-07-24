package main

import (
	"container/list"
	"fmt"
)

func main() {
	//new()创建空list
	//mylist:=list.New()
	//	//fmt.Println(//mylist,//mylist.Len())
	//	//fmt.Printf("%p",//mylist)


	mylist:=list.New()

	//添加元素
	//末尾添加
	mylist.PushBack("B")
	//头添加
	mylist.PushFront("A")
	//在第一个元素后面添加元素
	mylist.InsertAfter("a",mylist.Front())
	//在最后一个元素后面添加元素
	mylist.InsertBefore("b",mylist.Back())
	fmt.Println(mylist)

	//遍历显示元素内容
	for e:=mylist.Front();e!=nil;e=e.Next(){
		fmt.Print(e.Value," ")
	}
fmt.Println("")
	//获取头尾元素
	fmt.Println(mylist.Back().Value)
	fmt.Println(mylist.Front().Value)
	fmt.Println("***************************")
	//获取中间元素
	n:=2//获取的是第三个元素
	first:=mylist.Front()
	for i:=0;i<n;i++{
		first=first.Next()
	}
	fmt.Println(first.Value)

	fmt.Println("*********获取中间第2个元素***********")
	for e:=mylist.Front();e!=nil;e=e.Next(){
		fmt.Print(e.Value," ")
	}
	fmt.Println("")
	n=2
	var curr *list.Element
	if n>0 && n<mylist.Len(){
		if n==1{
			curr=mylist.Front()
		}else if n==mylist.Len(){
			curr=mylist.Back()
		}else{
			curr=mylist.Front()
			for i:=1;i<n ;i++  {
				curr=curr.Next()
			}
			fmt.Println("中间元素",n,curr.Value)
		}

	}else{
		fmt.Println("n的数值不对，超出list长度。")
	}

	//移动元素的顺序

	fmt.Println("*************移动元素***************")
	for e:=mylist.Front();e!=nil;e=e.Next(){
		fmt.Print(e.Value," ")
	}
	fmt.Println("")
	//把第一个移动到最后一个元素的后面
	//mylist.MoveToBack(mylist.Front())
	//把最后一个元素移动到最前面
	//mylist.MoveToFront(mylist.Back())
	//把第一个参数移动到第二个参数的后面。
	//mylist.MoveAfter(mylist.Front(),mylist.Back())
	//把第一个参数移动到第二个参数的前面。
	mylist.MoveBefore(mylist.Front(),mylist.Back())


	for e:=mylist.Front();e!=nil;e=e.Next(){
		fmt.Print(e.Value," ")
	}
	fmt.Println("")

	fmt.Println("**********删除元素**********")
	for e:=mylist.Front();e!=nil;e=e.Next(){
		fmt.Print(e.Value," ")
	}
	fmt.Println("")
	//删除
	mylist.Remove(mylist.Back())
	for e:=mylist.Front();e!=nil;e=e.Next(){
		fmt.Print(e.Value," ")
	}
	fmt.Println("")


}
