package main

import (
	"fmt"
	"unsafe"
)

//数据类型  整数
func main() {
	//test  范围 int -128-127
	//var j int8 =-129   //wrong
	//fmt.Println(j)

	//range 0-255
	var k uint16 =255
	fmt.Println(k)

	//int uint rune byte
	var a int =8900
	fmt.Println(a)
	var b uint =1
	var c byte =255
	fmt.Println(b,c)

	var n1 =100
	fmt.Printf("%T",n1)

	var  n2 int64 =10
	fmt.Printf("%T,%d",n1,unsafe.Sizeof(n2))
}
