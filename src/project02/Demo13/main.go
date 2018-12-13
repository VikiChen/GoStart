package main

import "fmt"

//数据类型与String类型 相互转化
func main() {
	var num int =10
	fmt.Println(num)
	prt :=&num
	fmt.Println(prt)
	*prt=20
	fmt.Println(num)
	}
