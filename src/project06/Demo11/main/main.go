package main

import "fmt"

//斐波那契数
func main() {
	fmt.Println(fub(10))
}

func fub(n int) ([]uint64)  {
	fbnSlice :=make([]uint64,n)
	fbnSlice[0]=1
	fbnSlice[1]=1
	for i:=2;i<n;i++{
		fbnSlice[i]=fbnSlice[i-1]+fbnSlice[i-2]
	}
	return fbnSlice
}