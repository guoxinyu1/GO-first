package main

import (
	"database/sql"
	"fmt"
	//此处不要忘记引入驱动
	_"github.com/go-sql-driver/mysql"
)

func main() {
	//1.打开链接
	db,err:=sql.Open("mysql","root:123456@tcp(localhost:3306)/first")
	db.Ping()//真正的打开
	defer func() {
		if db!=nil {
			db.Close()

		}
	}()
	if err !=nil{
		fmt.Println("数据库连接失败。")
		return
	}

	//2.预处理sql，？表示占位符
	stmt,err:=db.Prepare("insert  into people values (default ,?,?)")
	defer func() {
		if stmt!=nil {
			stmt.Close()

		}
	}()
	if err !=nil{
		fmt.Println("预处理失败。")
		return
	}
	//参数和占位符相对应
	result,err:=stmt.Exec("张三","海淀")
	//3.获取结果
	count,err:=result.RowsAffected()
	if err !=nil{
		fmt.Println("结果获取失败。")
		return
	}
	if count>0{
		fmt.Println("新增成功。")
	}else{
		fmt.Println("新增失败")
	}
	//4.关闭
	//result.close
	db.Close()
	stmt.Close()
}
