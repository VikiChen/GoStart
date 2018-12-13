package main

import (
	"fmt"
	"unsafe"
)

//数据类型  小数类型
func main() {

	var price float32 =89.123
	fmt.Println(price)

	var num1 float32=-0.0089
	var num2 float32=-878782.2
	fmt.Println(num1)
	fmt.Println(num2)

	//精度损失
	var num3 float32= -129.00000897
	var num4 float64= -129.00000897
	fmt.Println(num3)
	fmt.Println(num4)

	num5 :=1.1
	fmt.Printf("%T,%d",num5,unsafe.Sizeof(num5))


	num7 :=.123  //=>0.123
	fmt.Println(num7)

	//科学计数法
	num8 := 5.234e2
	fmt.Println(num8)

}
