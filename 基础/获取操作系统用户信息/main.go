package main

import (
	"fmt"
	"os/user"
)

func main() {
	u,error:=user.Current()
	if error!=nil{
		fmt.Println(error)
		return
	}
	fmt.Println(u.Uid)
	fmt.Println(u.Gid)
	fmt.Println(u.Name)
	fmt.Println(u.Username)
	fmt.Println(u.HomeDir)


	
}
