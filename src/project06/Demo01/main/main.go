package main

import "fmt"

//数组
func main() {
	var hens [6]float64
	hens[0]=2
	hens[1]=2
	hens[2]=2
	hens[3]=2

	totalweight :=0.0
	for i := 0; i< len(hens);i++  {
		totalweight+=hens[i]
	}


	fmt.Println(totalweight)
}