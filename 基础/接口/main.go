package main

import "fmt"

//定义接口关键字是interface
type Live interface {
	run(run int)
	eat()
}
//定义结构体准备实现接口
type People struct {
	name string
}
//结构体实现接口，重写接口内的所有方法才行
func (p People) run(run int)  {
	fmt.Println(p.name,"正在跑步，跑了：",run,"米")

}
func (p *People) eat() {
	fmt.Println(p.name,"正在吃饭。")

}

func main() {

	//实例化结构体People来调用接口
	p:=new(People)
	p.name="郭新宇"
	p.eat()
	p.run(5)

	
}
