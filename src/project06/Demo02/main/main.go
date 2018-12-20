package main

import "fmt"

//数组 初始化方式
func main() {
	var numArr01 [3]int =[3]int{1,2,3}
	fmt.Println(numArr01)

	var numArr02  =[3]int{1,2,3}
	fmt.Println(numArr02)

	var numArr03  =[...]int{1,2,3}
	fmt.Println(numArr03)

	var numArr04  =[...]int{1:800,2:200,0:3}
	fmt.Println(numArr04)

	str05:=[...]string{1:"tom",0:"jack",2:"mart",3:"mart"}
	fmt.Println(str05)
}