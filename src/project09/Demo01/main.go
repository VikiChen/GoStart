package main

import "fmt"

type stu struct {
	Name string
	Age int
	Scores [3]float64
	ptr *int
	slice []int
	map1 map[string]string
}

//结构体
func main() {
	var stu1 stu
	fmt.Println(stu1)
	fmt.Println("\n")
	stu1.slice=make([]int,2)
	stu1.slice[1]=2
	fmt.Println(stu1)
}
