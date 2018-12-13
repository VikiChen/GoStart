package main

import "fmt"

//数据类型转换  必须强制转换
func main() {

	var i int=100
	var n1 float32=float32(i)
	var n2 int32=int32(i)
	var n3 int64=int64(i)   //低精度到高精度 也需要强制转化

	fmt.Printf("%v,%v",i,n1)
	fmt.Println()
	fmt.Printf("%v,%v",i,n2)
	fmt.Println()
	fmt.Printf("%v,%v",i,n3)
	fmt.Println()

	var num1 int64=999999
	var num2 int8=int8(num1)   //会有溢出处理
	fmt.Println(num2)
}
