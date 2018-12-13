package main

import "fmt"

//数据类型与String类型 相互转化
func main() {
	var i ="sdwd"
	fmt.Println(&i)

	//*int 指向int类型的 指针
	var ptr *string =&i
	fmt.Println(ptr)
	fmt.Println(&ptr)
	fmt.Println(*ptr)

	}
