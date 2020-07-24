package main

import "fmt"

func main() {
	const a string  ="gxy"
	const b  =1 //int
	const c  =1.5 //float64
	fmt.Printf("a的类型是:%T\nb的类型是:%T\nc的类型是:%T\n",a,b,c)
	//常量可以直接进行数值运算，前提是要省去声明类型。
	fmt.Println(b+c)


	//定义常量组
	const(
		g=1
		x
		y
	)
	fmt.Println(g,x,y)

	//常量
	const(
		q=iota
		w
		e
	)
	fmt.Println(q,w,e)
	const(
		r=iota+8
		t
		u
	)
	fmt.Println(r,t,u)

	const(
		i=iota  //iota=0
		o       //iota=1
		p=3     //iota=2,p=3
		j       //iota=3,但是j=3
		k       //iota=4,k=3
		l=iota  //iota=5,l=5
	)
	fmt.Println(i,o,p,j,k,l)
	
}
