package main

import "fmt"
//二维 数组
func main() {
	var arr[4][6]int
	arr[1][2] =2
	fmt.Println(arr)


	for i,v :=range arr{
		fmt.Printf("%v,%v",i,v)
	}
}
