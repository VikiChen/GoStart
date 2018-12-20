package main

import (
	"math/rand"
	"fmt"
	"time"
)

//数组 翻转
func main() {
	var arr [5]int
	rand.Seed(time.Now().UnixNano())
	for i:=0;i< len(arr);i++{
		arr[i]=rand.Intn(100)
	}
	fmt.Println(arr)
	for i:=0;i< len(arr)/2;i++{
		temp:=arr[i]
		arr[i]=arr[len(arr)-i-1]
		arr[len(arr)-i-1]=temp
	}

	fmt.Println(arr)

}
