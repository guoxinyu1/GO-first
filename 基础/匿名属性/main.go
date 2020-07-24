package main

import "fmt"

type People struct {
	string
	int
}
func main(){
	p:=new(People)
	p.string="ss"
	p.int=12
	fmt.Println(p)
}