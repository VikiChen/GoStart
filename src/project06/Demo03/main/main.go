package main

import "fmt"

//数组 变量
func main() {

	var arr =[]int{1,2,34,5,5,66,}
	for i:=0;i< len(arr);i++{
		fmt.Println(arr[i])
	}

	for index,value :=range arr{
		fmt.Printf("%v,%v\n",index,value)
	}
}