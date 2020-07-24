package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type People struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}
type Peoples struct {
	XMLName xml.Name `xml:"peoples"`
	Version string		`xml:"version,attr"`
	//继承
	People []People     `xml:"people"`//反射获取标记，反射就是动态获取和操作结构体的属性。
}

func main() {
	//new新建结构体必须是指针
	peo := new(Peoples)
	//读取文件,返回的是一个字节切片和error
	b, _ := ioutil.ReadFile("D:/go_moudle/xml文件/people.xml")//反射
	//fmt.Println("读取xml文件：",string(b))
	//使用Unmarshal可以直接把XML字节切片数据转换为结构体。
	//此时也就是将people.xml里面的数据“<name>张三</name>”等
	//直接转换成结构体(多个变量)。
	xml.Unmarshal(b,peo)
	fmt.Println("将xml文件内容解析成结构体：\n",peo)

}
