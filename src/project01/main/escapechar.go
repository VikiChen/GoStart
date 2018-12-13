package main

import "fmt"

func main() {
	//转义字符
	fmt.Println("tom\tjack")
	fmt.Println("=================================")
	fmt.Println("tom\\jack")
	fmt.Println("=================================")
	fmt.Println("tom\"jack")
	fmt.Println("=================================")
	fmt.Println("tom\njack")
	fmt.Println("=================================")
	fmt.Println("tom\rjack")
	fmt.Println("==============test===============")
	fmt.Println("Name\tNative\tAge\nViki\tShanghai\t24")
}
