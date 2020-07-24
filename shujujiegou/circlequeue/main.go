package main

import (
	"errors"
	"fmt"
	"os"
)

//使用一个结构体管理环形数组
type CircleQueue struct {
	maxsize int
	array   [3]int
	tail    int //尾
	head    int //头
}

//入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("队列已满")
	}
	//分析出this.tail在队列尾部，但是包含最后的元素
	this.array[this.tail]=val
	this.tail=(this.tail+1)%this.maxsize
	return
}
//出队列
func (this *CircleQueue) Pop() (val int,err error) {
	if this.IsEmpty() {
		return 0,errors.New("队列为空！")
	}
	//分析出this.tail在队列尾部，但是包含最后的元素
	val=this.array[this.head]
	this.head=(this.head+1)%this.maxsize
	return
}

//队列为满
func (this *CircleQueue) IsFull() bool {

	return (this.tail+1)%this.maxsize == this.head
}

//队列为空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head

}
//取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	//关键算法。
	return (this.tail+this.maxsize-this.head)%this.maxsize

}
//显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下：")
	//取出当前元素有多少个元素。
	size:=this.Size()
	if size==0{
		fmt.Println("队列为空！")
	}
	//设计一个辅助变量指向head
	temphead:=this.head
	for i:=0;i<size;i++{
		fmt.Printf("arr[%d]:%d\t",temphead,this.array[temphead])
		temphead=(temphead+1)%this.maxsize
	}
	fmt.Println()


}

func main() {
	queue := &CircleQueue{
		maxsize: 3,
		head:   0,
		tail:    0,
	}
	//构建菜单用for+switch
	var key string
	var val int
	for {
		fmt.Println("1.输入add 表示添加数据到队列。")
		fmt.Println("2.输入show 表示显示队列数据。")
		fmt.Println("3.输入get 表示获取数据到队列。")
		fmt.Println("4.输入exit 表示退出。")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入一个数字插入到队列：")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())

			} else {

				fmt.Println("添加成功！")
			}

		case "show":
			queue.ListQueue()
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println("从队列取出数据失败！", err)
			} else {
				fmt.Println("从队列取出的数据是：", val)
			}
		case "exit":
			os.Exit(0)

		}
	}
}
