package main

import "fmt"

type People struct {
	Name string
	Weight int
}

func (p *People /*接收者*/) run()  {
	fmt.Println(p.Name,"正在跑步，当前体重为：",p.Weight)
	p.Weight-=1
	//fmt.Println("形参体重发生变化：",p.Weight)

}
func main() {
	pe0:=People{"gxy",100}//传递的是值副本
	pe0.run()//调用者
	fmt.Println(pe0.Name,"跑步完的体重：",pe0.Weight)
}
