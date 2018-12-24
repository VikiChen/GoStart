package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	//file ,err:=os.Open("d:/1.txt")
	//if err!=nil{
	//	fmt.Println("open file err=",err)
	//}
	content, err := ioutil.ReadFile("d:/1.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Println(string(content))
}