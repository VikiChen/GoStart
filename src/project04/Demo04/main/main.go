package main

import "fmt"

func sum(n1 ,n2 int)  int{
	defer fmt.Println("sum=n1",n1)
	defer fmt.Println("sum=n2",n2)
	res:=n1+n2
	fmt.Println("sum=res",res)
	return res
}

//defer
func main() {
	res:=sum(10,20)
	fmt.Println("main=res",res)

}



