package main

import "fmt"

//创建一个结构体来表示链表的一个节点。
type student struct {
	id   int
	name string
	next *student //指向下一个节点
}

//编写一种插入方法，在单链表的最后加入。[简单]
//思路：1.先找到该链表的最后一个节点。
//		2.创建一个辅助节点将头节点赋值给他，让他不断指向下一个节点来完成操作。
func InsertStudent(head *student, newstudent *student) {
	temp := head
	for {
		if temp.next == nil { //代表的是链表的最后一个节点
			break
		}
		temp = temp.next //一直循环操作，让temp一直指向下一个节点。
	}
	temp.next = newstudent //完成在最后一个节点插入
}

//显示链表
//思路：1.创建一个辅助节点帮忙
//		2.遍历链表temp=temp.next
func show(head *student) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空！")
		return
	}
	for {
		fmt.Printf(" %d:%s--> ", temp.next.id, temp.next.name)
		temp = temp.next
		if temp.next == nil {
			break
		}

	}
}

//根据节点的id属性的大小，从小到大完成插入到链表中。[常用]
//思路：1.找到适当的节点
//		2.创建一个辅助节点，让插入的节点和辅助节点的下一个节点的id属性比较。
//		3.让要插入的节点指向辅助节点的下一个节点，辅助节点指向插入节点，完成插入。

func InsertStudentbyid(head *student, newstudent *student) {
	temp := head
	flag := true
	for {
		if temp.next == nil { //代表的是链表的最后一个节点
			break
		} else if temp.next.id > newstudent.id {
			break //代表newstudent应在temp.next前插入链表
		} else if temp.next.id == newstudent.id {
			flag = false
			break
		}
		temp = temp.next //不要忘了让temp一直指向下一个节点。
	}
	if !flag {
		fmt.Println("对不起，已存在该id：", newstudent.id)
		return
	} else {
		newstudent.next = temp.next
		temp.next = newstudent //这两步顺序不能变否则会发生断链。
	}

}

//删除一个节点
//思路：1.创建一个辅助节点temp，根据方法传入的参数id和辅助节点的下一个节点的id作比较。
//		2.创建一个flag,通过改变flag的ture flase来执行删除。
//		3.删除操作就是：temp.next=temp.next.next-->next是节点的属性，存储的是节点类型的指针。

func delectnyid(head *student, id int) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.id == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next
	} else {
		fmt.Println("要删除的id不存在！")
	}

}

func updatebyid(head *student, id int, name string) {
	temp := head
	flag := false
	for {
		if temp.next == nil {
			break
		} else if temp.next.id == id {
			flag = true
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next.name = name
	} else {
		fmt.Println("要修改的id不存在！")
	}

}

//从头节点获取元素
func get(head *student) {
	temp := head
	if temp.next == nil {
		fmt.Println("链表为空！")
		return
	}
	temp.next = temp.next.next //每次跳过第一个节点达到取出元素。
}

func main() {
	//1.创建一个头节点，必不可少。
	head := &student{}
	//2.创建一个新节点/
	student1 := &student{
		id:   1,
		name: "张三",
		next: nil,
	}
	student2 := &student{
		id:   2,
		name: "kk",
		next: nil,
	}
	student4 := &student{
		id:   4,
		name: "李四",
		next: nil,
	}
	InsertStudent(head, student1)
	InsertStudent(head, student4)
	show(head)
	fmt.Println()
	InsertStudentbyid(head, student2)
	show(head)
	fmt.Println()
	delectnyid(head, 4)
	show(head)
	fmt.Println()
	updatebyid(head, 1, "郭新宇")
	show(head)
	fmt.Println()
	get(head)
	show(head)
}
