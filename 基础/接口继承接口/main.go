package main

import "fmt"
//定义接口
type Eat interface {
	eat()
}

//继承接口Eat
type Live interface {
	run(run int)
	Eat
}
//定义结构体准备实现接口
type People struct {
	name string
}
type Animal struct {

}
//结构体实现接口，重写接口内的所有方法才行
func (p People) run(run int)  {
	fmt.Println(p.name,"正在跑步，跑了：",run,"米")

}
func (p *People) eat() {
	fmt.Println(p.name,"正在吃饭。")

}

func (a *Animal) eat()  {
	fmt.Println("动物在吃饭。")

}

func main() {

	//实例化结构体People来调用接口
	p:=new(People)
	p.name="郭新宇"
	p.eat()
	p.run(5)

	a:=Animal{}
	a.eat()


}
