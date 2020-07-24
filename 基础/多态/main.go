package main

import "fmt"

type Live interface {
	run()
}

type People struct{}
type Animal struct{}

func (p *People) run() {
	fmt.Println("人在跑步。")

}
func (p *Animal) eat() {
	fmt.Println("动物在跑步。")

}

//多态的体现:接口变量作为方法参数。
func AllRun(live Live) {
	live.run()
}
func main() {
	var live Live = &People{} //var live Live=&Animal{}

	AllRun(live)
}
