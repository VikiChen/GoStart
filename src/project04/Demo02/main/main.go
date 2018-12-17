package main

import "fmt"

//闭包
func main() {
	f:=addUpper()

	fmt.Println(f(1))
	fmt.Println(f(2))
}

func addUpper() func(int) int {
	var n int =10
	return func(x int) int {
		n=n+x
		return n
	}
}
