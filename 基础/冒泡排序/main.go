package main

import "fmt"

func main() {
	arr :=[5]int{18,98,25,3,0}
	for i:=0;i<len(arr)-1;i++{
		for j:=0;j<len(arr)-1-i ;j++  {
			if arr[j]>arr[j+1] {
				arr[j],arr[j+1]=arr[j+1],arr[j]

			}

		}
	}
	fmt.Println(arr)
	
}
