package main

import "fmt"

//切片 变量
func main() {

	//常规for 循环遍历方式
	var arr =[5]int{10,20,20,343,354}
	slice :=arr[1:4]
	for i:=0;i< len(slice);i++{
		fmt.Println(slice[i])
	}

	for i,v :=range slice{
		fmt.Printf("%v,%v\n",i,v)
	}
}
