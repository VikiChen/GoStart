package main

import "fmt"

//切片 变量
func main() {

	//常规for 循环遍历方式
	var arr =[5]int{10,20,20,343,354}
	slice :=arr[1:4]
	slice=append(slice, 22,233)
	fmt.Println(slice)
}
