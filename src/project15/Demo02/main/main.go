package main

import "fmt"

func main() {
	var num int
	const tax int=1

	fmt.Println(num)
	fmt.Println(tax)
	const (
		a=iota
		b
		c
		d
	)
	fmt.Println(a,b,c,d)

}