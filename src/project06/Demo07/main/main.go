package main

import "fmt"

//切片 使用
func main() {
	var slice []float64 =make([]float64,5,10)
	//对于切片 ，必须make
	fmt.Println(slice)

	var stringSlice []string =[]string{"goi","dsd","sdsd"}
	fmt.Println(len(stringSlice))
	fmt.Println(stringSlice)
	fmt.Println(cap(stringSlice))
}
