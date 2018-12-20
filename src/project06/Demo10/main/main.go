package main

import "fmt"

//切片 string 和 slice
func main() {
	str :="sjdhkshkd:2"
	slice :=str[:]
	fmt.Println(slice)
	for i:=0;i< len(slice);i++{
		fmt.Printf("%c\n",slice[i])
	}
}
