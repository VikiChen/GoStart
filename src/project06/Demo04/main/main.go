package main

import "fmt"

//数组 变量
func main() {
	var arr [26]byte
	for i:=0;i<26;i++{
		arr[i]='A'+byte(i)
	}
	fmt.Printf("%c",arr)
}