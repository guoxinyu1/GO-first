package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	row int//行
	col int//列
	val int//值
}

func main() {
	fmt.Println("***************原二维数组:****************")
	//1.定义一个二维数组实现五子棋
	var arr [11][11]int
	arr[1][2] = 1 //代表白子
	arr[2][3] = 2 //代表黑子
	for _, v := range arr {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
	//2.将二维数组转成稀疏数组。想->算法。
	//思路：(1).遍历二维数组，当发现元素不为0时，创建一个Node结构体，
	//(2).将其放入到对应的切片中。
	var sparceArr []Node //存储稀疏数组
	for i, v := range arr {
		for j, v2 := range v {
			if v2 != 0 {
				valNode := Node{
					row: i,
					col: j,
					val: v2,
				}
				sparceArr = append(sparceArr, valNode)
			}
		}
		fmt.Println()
	}
	//3.存盘
	fmt.Println("***************稀疏二维数组:****************")
	filePath := "D:/wuqiqi.data"
	file, err := os.OpenFile(filePath, os.O_APPEND, 0666)
	if err!=nil{
		file,_=os.Create(filePath)
	}
	defer file.Close()
	for _, v := range sparceArr {
		fmt.Printf("%d %d %d \n", v.row, v.col, v.val)
		//将int转换为字符串切片。并以分割符“ ”，合并成字符串写入到文件中。
		str := []string{strconv.Itoa(v.row), strconv.Itoa(v.col), strconv.Itoa(v.val), "\n"}
		strs := strings.Join(str, " ")
		writer := bufio.NewWriter(file)
		writer.WriteString(strs)
		writer.Flush()//刷新缓存。
	}

	//4.文件读取，恢复到原来的二维数组
	fmt.Println("***************还原原二维数组:****************")

	var arr2 [11][11]int
	file2, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return
	}
	defer file2.Close()
	//遍历读取每一行的字符串，以'\n'为结束符读入一行
	rd := bufio.NewReader(file2)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		s1 := strings.Split(line, " ")//将字符串分割成切片，才能使用索引。
		a, _ := strconv.Atoi(s1[0])
		b, _ := strconv.Atoi(s1[1])
		c, _ := strconv.Atoi(s1[2])
		arr2[a][b] = c
	}
	//遍历二维数组。
	for _, v := range arr2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}