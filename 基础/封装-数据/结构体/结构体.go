package 结构体
//	People结构体可以被跨包访问到
type People struct {
	//结构体内的属性首字母小写，不能直接被访问到。
	//需要通过方法来进行访问。
	//封装数据，方法来掉
	name string
	age int

}
//func (变量名 结构体类型) 方法名(参数名 参数类型) 返回值列表{}
//设置属性name
func (p *People) SetName(name string)  {
	p.name=name
}
//获取属性
func (p *People) GetName() string{
	return p.name
}

//设置属性age
//在一个方法中参数是另一个结构体的变量。--->依赖关系。
func (p *People) SetAge(age int)  {
	if age<0 || age>150 {
		p.age=0

	}else {
		p.age=age
	}

}
//获取age

func (p *People) GetAge() int{
	return p.age
}