package main

import (
	"fmt"
	"unsafe"
)

//数据类型  boolen
func main() {


	var b bool =true
	fmt.Println(b)
	//注意事项
	//1.占用空间1字节
	fmt.Printf("%d",unsafe.Sizeof(b))
	//2. 只有true 和false

}
