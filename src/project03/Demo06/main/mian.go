package main

import "fmt"

func main() {

	label1:
	for i:=1 ;i<100;i++ {
		for j:=1 ;j<100;j++ {
			fmt.Println(i," ",j)
			break label1
		}
	}

}
