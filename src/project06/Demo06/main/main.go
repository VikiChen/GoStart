package main

import "fmt"

//切片
func main() {

	var intArr [5]int=[...]int{1,2,3,4,5}
	slice :=intArr[1:3]  //不包含3
	fmt.Println(intArr)
	fmt.Println(len(intArr))
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
}
