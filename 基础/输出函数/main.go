package main

import "fmt"

func main() {

	//fmt.Fprintln(os.Stdout,"内容1","内容2")
	//fmt.Fprintln(os.Stdout,"内容3","内容4")
	//fmt.Fprint(os.Stdout,"内容1","内容2")
	a:=fmt.Sprintln("a","b")
	fmt.Println(a)
	b:=fmt.Sprint("1","2","3")
	fmt.Println(b)

}
