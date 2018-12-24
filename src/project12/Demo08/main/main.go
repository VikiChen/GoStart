package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file1Path :="d:/abd.txt"
	file2Path:="d:kkk.txt"
	content, e := ioutil.ReadFile(file1Path)
	if e!=nil{
		fmt.Println("readfileerr",e)
		return
	}
	err := ioutil.WriteFile(file2Path, content, 0666)
	if err!=nil{
		fmt.Println("writer err",err)
	}

}