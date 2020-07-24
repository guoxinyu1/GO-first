package main

import "fmt"

type People struct {
	name string
	age int
}
type Teacher struct {
	//包含另一个结构体作为属性。
	peo People
	classroom string
}

func main() {
	teacher:=new(Teacher)
	teacher.classroom="1班"
	teacher.peo.name="gxy"
	teacher.peo.age=18
	fmt.Println(teacher)
	
}
