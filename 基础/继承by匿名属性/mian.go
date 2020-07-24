package main

import "fmt"
//父结构体
type People struct {
	name string
	age int
}
//子结构体
type Teacher struct {
	//匿名属性，完成继承People。
	People
	classroom string
}

func main() {
	teacher:=new(Teacher)
	teacher.classroom="1班"
	teacher.name="gxy"
	teacher.age=18
	fmt.Println(teacher)

}
