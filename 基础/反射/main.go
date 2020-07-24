package main

import (
	"fmt"
	"reflect"
)

type People struct {
	Id    int `xml:"id"`//标记
	Name string `xml:"adress"`
}

func main() {
	peo:=People{1,"郭新宇"}

	//反射:动态操作结构体的属性。
	v := reflect.ValueOf(peo)
	//获取结构体属性个数
	fmt.Println(v.NumField())
	//通过索引获得目标属性
	fmt.Println(v.FieldByIndex([]int{1}))

	//通过变量 反射 获取结构体内的属性的值。
	a:="name"
	b:="id"
	fmt.Println(v.FieldByName(a))
	fmt.Println(v.FieldByName(b))

	//设置 结构体属性的值
	peo2:=new(People)
	//peo2.Name="L"
	//peo2.Id=2
	v1:=reflect.ValueOf(peo2).Elem()
	fmt.Println(v1.FieldByName("id").CanSet())
	v1.FieldByName("Id").SetInt(123)
	v1.FieldByName("Name").SetString("lisi")
	fmt.Println(peo2)

	//获取标记
	t:=reflect.TypeOf(People{})
	name,_:=t.FieldByName("Name")
	fmt.Println(name.Tag)//获取完整标记
	fmt.Println(name.Tag.Get("xml"))//获取标记中的内容。
}
