package main

import "fmt"

func main() {

	var p = 10
	for i:=1;i<p ;i++  {
		for k:=1;k<=p-i;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=2*i-1 ;j++ {
			if j==1 || j==2*i-1{
				fmt.Print("*")
			}else if i==p-1{
				fmt.Print("*")
			} else{
				fmt.Print(" ")
			}

		}
		fmt.Println("\n")

	}
}
