package main

import (
	"errors"
	"fmt"
	"os"
)

/*此代码实现了基本队列结构，但是没有实现空间复用，它是单向的。*/

//使用一个结构体管理队列
type queue struct {
	maxsize int
	array   [3]int
	front   int
	rear    int
}

//添加数据到队列
func (this *queue) add(val int) (err error) {
	//先判断队列是否已满
	if this.rear == this.maxsize-1 { //重要提示rear包含队列尾部最后一个元素。
		return errors.New("队列已满！")
	}
	this.rear++
	this.array[this.rear] = val
	return
}

//从队列中取值，只能从前端取值
func (this *queue) GetQueue() (val int, err error) {
	//判断队列是否为空
	if this.front == this.rear {
		return -1, errors.New("queue empty!")
	}
	this.front++
	val = this.array[this.front]
	return val, err
}

//显示队列,先找到队首，从队首遍历到最后一个元素
//重要提示front指向队列第一个元素的前面，不包含第一个元素。
//第一个元素:front+1
func (this *queue) showQueue() {
	fmt.Println("队列当前情况时：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]:%d", i, this.array[i])
	}
	fmt.Println()
}
func main() {
	queue := &queue{
		maxsize: 3,
		front:   -1,
		rear:    -1,
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
			err := queue.add(val)
			if err != nil {
				fmt.Println(err.Error())

			} else {

				fmt.Println("添加成功！")
			}

		case "show":
			queue.showQueue()
		case "get":
			val, err := queue.GetQueue()
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
